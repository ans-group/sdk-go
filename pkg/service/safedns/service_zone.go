package safedns

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	return connection.InvokeRequestAll(s.GetZonesPaginated, parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error) {
	body, err := connection.Get[[]Zone](s.connection, "/safedns/v1/zones", parameters)
	return connection.NewPaginated(body, parameters, s.GetZonesPaginated), err
}

// GetZone retrieves a single zone by name
func (s *Service) GetZone(zoneName string) (Zone, error) {
	if zoneName == "" {
		return Zone{}, fmt.Errorf("invalid zone name")
	}
	body, err := connection.Get[Zone](s.connection, fmt.Sprintf("/safedns/v1/zones/%s", zoneName), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
	return body.Data, err
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
	if zoneName == "" {
		return fmt.Errorf("invalid zone name")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/safedns/v1/zones/%s", zoneName), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
}

// GetZoneRecords retrieves a list of records
func (s *Service) GetZoneRecords(zoneName string, parameters connection.APIRequestParameters) ([]Record, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetZoneRecordsPaginated(zoneName, p)
	}, parameters)
}

// GetZoneRecordsPaginated retrieves a paginated list of zones
func (s *Service) GetZoneRecordsPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	if zoneName == "" {
		return nil, fmt.Errorf("invalid zone name")
	}
	body, err := connection.Get[[]Record](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records", zoneName), parameters, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetZoneRecordsPaginated(zoneName, p)
	}), err
}

// GetZoneRecord retrieves a single zone record by ID
func (s *Service) GetZoneRecord(zoneName string, recordID int) (Record, error) {
	if zoneName == "" {
		return Record{}, fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return Record{}, fmt.Errorf("invalid record id")
	}
	body, err := connection.Get[Record](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}))
	return body.Data, err
}

// CreateZoneRecord creates a new SafeDNS zone record
func (s *Service) CreateZoneRecord(zoneName string, req CreateRecordRequest) (int, error) {
	if zoneName == "" {
		return 0, fmt.Errorf("invalid zone name")
	}
	body, err := connection.Post[Record](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records", zoneName), &req, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
	return body.Data.ID, err
}

// UpdateZoneRecord updates a SafeDNS zone record
func (s *Service) UpdateZoneRecord(zoneName string, record Record) (int, error) {
	if zoneName == "" {
		return 0, fmt.Errorf("invalid zone name")
	}
	if record.ID < 1 {
		return 0, fmt.Errorf("invalid record id")
	}
	body, err := connection.Put[Record](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, record.ID), &record, connection.NotFoundResponseHandler(&ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: record.ID}))
	return body.Data.ID, err
}

// PatchZoneRecord patches a SafeDNS zone record
func (s *Service) PatchZoneRecord(zoneName string, recordID int, patch PatchRecordRequest) (int, error) {
	if zoneName == "" {
		return 0, fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return 0, fmt.Errorf("invalid record id")
	}
	body, err := connection.Put[Record](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), &patch, connection.NotFoundResponseHandler(&ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}))
	return body.Data.ID, err
}

// DeleteZoneRecord removes a SafeDNS zone record
func (s *Service) DeleteZoneRecord(zoneName string, recordID int) error {
	if zoneName == "" {
		return fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return fmt.Errorf("invalid record id")
	}
	return connection.DeleteRaw(s.connection, fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), nil, &connection.APIResponseBody{}, connection.NotFoundResponseHandler(&ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}))
}

// GetZoneNotes retrieves a list of notes
func (s *Service) GetZoneNotes(zoneName string, parameters connection.APIRequestParameters) ([]Note, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Note], error) {
		return s.GetZoneNotesPaginated(zoneName, p)
	}, parameters)
}

// GetZoneNotesPaginated retrieves a paginated list of zones
func (s *Service) GetZoneNotesPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Note], error) {
	if zoneName == "" {
		return nil, fmt.Errorf("invalid zone name")
	}
	body, err := connection.Get[[]Note](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/notes", zoneName), parameters, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Note], error) {
		return s.GetZoneNotesPaginated(zoneName, p)
	}), err
}

// GetZoneNote retrieves a single zone note by ID
func (s *Service) GetZoneNote(zoneName string, noteID int) (Note, error) {
	if zoneName == "" {
		return Note{}, fmt.Errorf("invalid zone name")
	}
	if noteID < 1 {
		return Note{}, fmt.Errorf("invalid note id")
	}
	body, err := connection.Get[Note](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/notes/%d", zoneName, noteID), connection.APIRequestParameters{}, connection.NotFoundResponseHandler(&ZoneNoteNotFoundError{ZoneName: zoneName, NoteID: noteID}))
	return body.Data, err
}

// CreateZoneNote creates a new SafeDNS zone note
func (s *Service) CreateZoneNote(zoneName string, req CreateNoteRequest) (int, error) {
	if zoneName == "" {
		return 0, fmt.Errorf("invalid zone name")
	}
	body, err := connection.Post[Note](s.connection, fmt.Sprintf("/safedns/v1/zones/%s/notes", zoneName), &req, connection.NotFoundResponseHandler(&ZoneNotFoundError{ZoneName: zoneName}))
	return body.Data.ID, err
}
