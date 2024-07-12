package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial instantiates a new ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial and sets the default values.
func NewItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial()(*ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial) {
    m := &ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterial) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemAttestationsItemWithSubject_digestGetResponse_attestations_bundle_verificationMaterialable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
