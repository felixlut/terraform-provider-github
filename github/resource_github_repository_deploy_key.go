package github

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryDeployKeyCreate,
		Read:   resourceGithubRepositoryDeployKeyRead,
		Delete: resourceGithubRepositoryDeployKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		// Deploy keys are defined immutable in the API. Updating results in force new.
		Schema: map[string]*schema.Schema{
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressDeployKeyDiff,
				Description:      "A SSH key.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "A boolean qualifying the key to be either read only or read/write.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the GitHub repository.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A title.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryDeployKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	repoName := d.Get("repository").(string)
	key := d.Get("key").(string)
	title := d.Get("title").(string)
	readOnly := d.Get("read_only").(bool)
	owner := meta.(*Owner).name
	ctx := context.Background()

	resultKey, _, err := client.Repositories.CreateKey(ctx, owner, repoName, &github.Key{
		Key:      github.String(key),
		Title:    github.String(title),
		ReadOnly: github.Bool(readOnly),
	})

	if err != nil {
		return err
	}

	id := strconv.FormatInt(resultKey.GetID(), 10)

	d.SetId(buildTwoPartID(repoName, id))

	return resourceGithubRepositoryDeployKeyRead(d, meta)
}

func resourceGithubRepositoryDeployKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, idString, err := parseTwoPartID(d.Id(), "repository", "ID")
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	key, resp, err := client.Repositories.GetKey(ctx, owner, repoName, id)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository deploy key %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return err
	}
	if err = d.Set("key", key.GetKey()); err != nil {
		return err
	}
	if err = d.Set("read_only", key.GetReadOnly()); err != nil {
		return err
	}
	if err = d.Set("repository", repoName); err != nil {
		return err
	}
	if err = d.Set("title", key.GetTitle()); err != nil {
		return err
	}

	return nil
}

func resourceGithubRepositoryDeployKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, idString, err := parseTwoPartID(d.Id(), "repository", "ID")
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err = client.Repositories.DeleteKey(ctx, owner, repoName, id)
	if err != nil {
		return err
	}

	return err
}

func suppressDeployKeyDiff(k, oldV, newV string, d *schema.ResourceData) bool {
	newV = strings.TrimSpace(newV)
	keyRe := regexp.MustCompile(`^([a-z0-9-]+ [^\s]+)( [^\s]+)?$`)
	newTrimmed := keyRe.ReplaceAllString(newV, "$1")

	return oldV == newTrimmed
}
