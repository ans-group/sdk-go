package safedns

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/internal/resource"
)

func (s *Service) zoneRes() *resource.Resource[Zone, string] {
	return resource.NewStringResourceWithIdentifier[Zone](s.connection, "/safedns/v1/zones", "zone", "name",
		func(id string) error { return &ZoneNotFoundError{ZoneName: id} })
}

func (s *Service) zoneRecordRes() *resource.SubResource[Record, string, int] {
	return resource.NewStringIntSubResource[Record](
		s.connection,
		func(zoneName string) string { return fmt.Sprintf("/safedns/v1/zones/%s/records", zoneName) },
		"zone", "name", func(zoneName string) error { return &ZoneNotFoundError{ZoneName: zoneName} },
		"record", "id", func(zoneName string, recordID int) error {
			return &ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}
		},
	)
}

func (s *Service) zoneNoteRes() *resource.SubResource[Note, string, int] {
	return resource.NewStringIntSubResource[Note](
		s.connection,
		func(zoneName string) string { return fmt.Sprintf("/safedns/v1/zones/%s/notes", zoneName) },
		"zone", "name", func(zoneName string) error { return &ZoneNotFoundError{ZoneName: zoneName} },
		"note", "id", func(zoneName string, noteID int) error {
			return &ZoneNoteNotFoundError{ZoneName: zoneName, NoteID: noteID}
		},
	)
}

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	return s.zoneRes().List(parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error) {
	return s.zoneRes().ListPaginated(parameters)
}

// GetZone retrieves a single zone by name
func (s *Service) GetZone(zoneName string) (Zone, error) {
	return s.zoneRes().Get(zoneName)
}

// CreateZone creates a new SafeDNS zone
func (s *Service) CreateZone(req CreateZoneRequest) error {
	return connection.PostRaw(s.connection, "/safedns/v1/zones", &req, &connection.APIResponseBody{})
}

// PatchZone patches a SafeDNS zone
func (s *Service) PatchZone(zoneName string, req PatchZoneRequest) error {
	if zoneName == "" {
		return fmt.Errorf("invalid zone name")
	}
	return connection.PatchRaw(s.connection, fmt.Sprintf("/safedns/v1/zones/%s", zoneName), &req, &connection.APIResponseBody{})
}

// DeleteZone removes a SafeDNS zone
func (s *Service) DeleteZone(zoneName string) error {
	return s.zoneRes().Delete(zoneName)
}

// GetZoneRecords retrieves a list of records
func (s *Service) GetZoneRecords(zoneName string, parameters connection.APIRequestParameters) ([]Record, error) {
	return s.zoneRecordRes().List(zoneName, parameters)
}

// GetZoneRecordsPaginated retrieves a paginated list of zones
func (s *Service) GetZoneRecordsPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	return s.zoneRecordRes().ListPaginated(zoneName, parameters)
}

// GetZoneRecord retrieves a single zone record by ID
func (s *Service) GetZoneRecord(zoneName string, recordID int) (Record, error) {
	return s.zoneRecordRes().Get(zoneName, recordID)
}

// CreateZoneRecord creates a new SafeDNS zone record
func (s *Service) CreateZoneRecord(zoneName string, req CreateRecordRequest) (int, error) {
	data, err := s.zoneRecordRes().Create(zoneName, &req)
	return data.ID, err
}

// UpdateZoneRecord updates a SafeDNS zone record
func (s *Service) UpdateZoneRecord(zoneName string, record Record) (int, error) {
	data, err := s.zoneRecordRes().Update(zoneName, record.ID, &record)
	return data.ID, err
}

// PatchZoneRecord patches a SafeDNS zone record
func (s *Service) PatchZoneRecord(zoneName string, recordID int, patch PatchRecordRequest) (int, error) {
	data, err := s.zoneRecordRes().Update(zoneName, recordID, &patch)
	return data.ID, err
}

// DeleteZoneRecord removes a SafeDNS zone record
func (s *Service) DeleteZoneRecord(zoneName string, recordID int) error {
	return s.zoneRecordRes().Delete(zoneName, recordID)
}

// GetZoneNotes retrieves a list of notes
func (s *Service) GetZoneNotes(zoneName string, parameters connection.APIRequestParameters) ([]Note, error) {
	return s.zoneNoteRes().List(zoneName, parameters)
}

// GetZoneNotesPaginated retrieves a paginated list of zones
func (s *Service) GetZoneNotesPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Note], error) {
	return s.zoneNoteRes().ListPaginated(zoneName, parameters)
}

// GetZoneNote retrieves a single zone note by ID
func (s *Service) GetZoneNote(zoneName string, noteID int) (Note, error) {
	return s.zoneNoteRes().Get(zoneName, noteID)
}

// CreateZoneNote creates a new SafeDNS zone note
func (s *Service) CreateZoneNote(zoneName string, req CreateNoteRequest) (int, error) {
	data, err := s.zoneNoteRes().Create(zoneName, &req)
	return data.ID, err
}
