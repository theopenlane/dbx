// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package dbxclient

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/theopenlane/dbx/pkg/enums"
)

type Dbxclient interface {
	CreateDatabase(ctx context.Context, input CreateDatabaseInput, interceptors ...clientv2.RequestInterceptor) (*CreateDatabase, error)
	DeleteDatabase(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*DeleteDatabase, error)
	GetAllDatabases(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllDatabases, error)
	GetDatabase(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*GetDatabase, error)
	GetGroup(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*GetGroup, error)
	GetAllGroups(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllGroups, error)
	CreateGroup(ctx context.Context, input CreateGroupInput, interceptors ...clientv2.RequestInterceptor) (*CreateGroup, error)
	DeleteGroup(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*DeleteGroup, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) Dbxclient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type CreateDatabase_CreateDatabase_Database struct {
	Dsn            string                 "json:\"dsn\" graphql:\"dsn\""
	Geo            *string                "json:\"geo,omitempty\" graphql:\"geo\""
	ID             string                 "json:\"id\" graphql:\"id\""
	Name           string                 "json:\"name\" graphql:\"name\""
	OrganizationID string                 "json:\"organizationID\" graphql:\"organizationID\""
	Provider       enums.DatabaseProvider "json:\"provider\" graphql:\"provider\""
	Status         enums.DatabaseStatus   "json:\"status\" graphql:\"status\""
}

func (t *CreateDatabase_CreateDatabase_Database) GetDsn() string {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return t.Dsn
}
func (t *CreateDatabase_CreateDatabase_Database) GetGeo() *string {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return t.Geo
}
func (t *CreateDatabase_CreateDatabase_Database) GetID() string {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return t.ID
}
func (t *CreateDatabase_CreateDatabase_Database) GetName() string {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return t.Name
}
func (t *CreateDatabase_CreateDatabase_Database) GetOrganizationID() string {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return t.OrganizationID
}
func (t *CreateDatabase_CreateDatabase_Database) GetProvider() *enums.DatabaseProvider {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return &t.Provider
}
func (t *CreateDatabase_CreateDatabase_Database) GetStatus() *enums.DatabaseStatus {
	if t == nil {
		t = &CreateDatabase_CreateDatabase_Database{}
	}
	return &t.Status
}

type CreateDatabase_CreateDatabase struct {
	Database CreateDatabase_CreateDatabase_Database "json:\"database\" graphql:\"database\""
}

func (t *CreateDatabase_CreateDatabase) GetDatabase() *CreateDatabase_CreateDatabase_Database {
	if t == nil {
		t = &CreateDatabase_CreateDatabase{}
	}
	return &t.Database
}

type DeleteDatabase_DeleteDatabase struct {
	DeletedID string "json:\"deletedID\" graphql:\"deletedID\""
}

func (t *DeleteDatabase_DeleteDatabase) GetDeletedID() string {
	if t == nil {
		t = &DeleteDatabase_DeleteDatabase{}
	}
	return t.DeletedID
}

type GetAllDatabases_Databases_Edges_Node struct {
	Dsn            string                 "json:\"dsn\" graphql:\"dsn\""
	Geo            *string                "json:\"geo,omitempty\" graphql:\"geo\""
	ID             string                 "json:\"id\" graphql:\"id\""
	Name           string                 "json:\"name\" graphql:\"name\""
	OrganizationID string                 "json:\"organizationID\" graphql:\"organizationID\""
	Provider       enums.DatabaseProvider "json:\"provider\" graphql:\"provider\""
	Status         enums.DatabaseStatus   "json:\"status\" graphql:\"status\""
}

func (t *GetAllDatabases_Databases_Edges_Node) GetDsn() string {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return t.Dsn
}
func (t *GetAllDatabases_Databases_Edges_Node) GetGeo() *string {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return t.Geo
}
func (t *GetAllDatabases_Databases_Edges_Node) GetID() string {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return t.ID
}
func (t *GetAllDatabases_Databases_Edges_Node) GetName() string {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return t.Name
}
func (t *GetAllDatabases_Databases_Edges_Node) GetOrganizationID() string {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return t.OrganizationID
}
func (t *GetAllDatabases_Databases_Edges_Node) GetProvider() *enums.DatabaseProvider {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return &t.Provider
}
func (t *GetAllDatabases_Databases_Edges_Node) GetStatus() *enums.DatabaseStatus {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges_Node{}
	}
	return &t.Status
}

type GetAllDatabases_Databases_Edges struct {
	Node *GetAllDatabases_Databases_Edges_Node "json:\"node,omitempty\" graphql:\"node\""
}

func (t *GetAllDatabases_Databases_Edges) GetNode() *GetAllDatabases_Databases_Edges_Node {
	if t == nil {
		t = &GetAllDatabases_Databases_Edges{}
	}
	return t.Node
}

type GetAllDatabases_Databases struct {
	Edges []*GetAllDatabases_Databases_Edges "json:\"edges,omitempty\" graphql:\"edges\""
}

func (t *GetAllDatabases_Databases) GetEdges() []*GetAllDatabases_Databases_Edges {
	if t == nil {
		t = &GetAllDatabases_Databases{}
	}
	return t.Edges
}

type GetDatabase_Database struct {
	Dsn            string                 "json:\"dsn\" graphql:\"dsn\""
	Geo            *string                "json:\"geo,omitempty\" graphql:\"geo\""
	ID             string                 "json:\"id\" graphql:\"id\""
	Name           string                 "json:\"name\" graphql:\"name\""
	OrganizationID string                 "json:\"organizationID\" graphql:\"organizationID\""
	Provider       enums.DatabaseProvider "json:\"provider\" graphql:\"provider\""
	Status         enums.DatabaseStatus   "json:\"status\" graphql:\"status\""
}

func (t *GetDatabase_Database) GetDsn() string {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return t.Dsn
}
func (t *GetDatabase_Database) GetGeo() *string {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return t.Geo
}
func (t *GetDatabase_Database) GetID() string {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return t.ID
}
func (t *GetDatabase_Database) GetName() string {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return t.Name
}
func (t *GetDatabase_Database) GetOrganizationID() string {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return t.OrganizationID
}
func (t *GetDatabase_Database) GetProvider() *enums.DatabaseProvider {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return &t.Provider
}
func (t *GetDatabase_Database) GetStatus() *enums.DatabaseStatus {
	if t == nil {
		t = &GetDatabase_Database{}
	}
	return &t.Status
}

type GetGroup_Group struct {
	Description     *string      "json:\"description,omitempty\" graphql:\"description\""
	ID              string       "json:\"id\" graphql:\"id\""
	Locations       []string     "json:\"locations,omitempty\" graphql:\"locations\""
	Name            string       "json:\"name\" graphql:\"name\""
	PrimaryLocation string       "json:\"primaryLocation\" graphql:\"primaryLocation\""
	Region          enums.Region "json:\"region\" graphql:\"region\""
}

func (t *GetGroup_Group) GetDescription() *string {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return t.Description
}
func (t *GetGroup_Group) GetID() string {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return t.ID
}
func (t *GetGroup_Group) GetLocations() []string {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return t.Locations
}
func (t *GetGroup_Group) GetName() string {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return t.Name
}
func (t *GetGroup_Group) GetPrimaryLocation() string {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return t.PrimaryLocation
}
func (t *GetGroup_Group) GetRegion() *enums.Region {
	if t == nil {
		t = &GetGroup_Group{}
	}
	return &t.Region
}

type GetAllGroups_Groups_Edges_Node struct {
	Description     *string      "json:\"description,omitempty\" graphql:\"description\""
	ID              string       "json:\"id\" graphql:\"id\""
	Locations       []string     "json:\"locations,omitempty\" graphql:\"locations\""
	Name            string       "json:\"name\" graphql:\"name\""
	PrimaryLocation string       "json:\"primaryLocation\" graphql:\"primaryLocation\""
	Region          enums.Region "json:\"region\" graphql:\"region\""
}

func (t *GetAllGroups_Groups_Edges_Node) GetDescription() *string {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return t.Description
}
func (t *GetAllGroups_Groups_Edges_Node) GetID() string {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return t.ID
}
func (t *GetAllGroups_Groups_Edges_Node) GetLocations() []string {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return t.Locations
}
func (t *GetAllGroups_Groups_Edges_Node) GetName() string {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return t.Name
}
func (t *GetAllGroups_Groups_Edges_Node) GetPrimaryLocation() string {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return t.PrimaryLocation
}
func (t *GetAllGroups_Groups_Edges_Node) GetRegion() *enums.Region {
	if t == nil {
		t = &GetAllGroups_Groups_Edges_Node{}
	}
	return &t.Region
}

type GetAllGroups_Groups_Edges struct {
	Node *GetAllGroups_Groups_Edges_Node "json:\"node,omitempty\" graphql:\"node\""
}

func (t *GetAllGroups_Groups_Edges) GetNode() *GetAllGroups_Groups_Edges_Node {
	if t == nil {
		t = &GetAllGroups_Groups_Edges{}
	}
	return t.Node
}

type GetAllGroups_Groups struct {
	Edges []*GetAllGroups_Groups_Edges "json:\"edges,omitempty\" graphql:\"edges\""
}

func (t *GetAllGroups_Groups) GetEdges() []*GetAllGroups_Groups_Edges {
	if t == nil {
		t = &GetAllGroups_Groups{}
	}
	return t.Edges
}

type CreateGroup_CreateGroup_Group struct {
	Description     *string      "json:\"description,omitempty\" graphql:\"description\""
	ID              string       "json:\"id\" graphql:\"id\""
	Locations       []string     "json:\"locations,omitempty\" graphql:\"locations\""
	Name            string       "json:\"name\" graphql:\"name\""
	PrimaryLocation string       "json:\"primaryLocation\" graphql:\"primaryLocation\""
	Region          enums.Region "json:\"region\" graphql:\"region\""
}

func (t *CreateGroup_CreateGroup_Group) GetDescription() *string {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return t.Description
}
func (t *CreateGroup_CreateGroup_Group) GetID() string {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return t.ID
}
func (t *CreateGroup_CreateGroup_Group) GetLocations() []string {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return t.Locations
}
func (t *CreateGroup_CreateGroup_Group) GetName() string {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return t.Name
}
func (t *CreateGroup_CreateGroup_Group) GetPrimaryLocation() string {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return t.PrimaryLocation
}
func (t *CreateGroup_CreateGroup_Group) GetRegion() *enums.Region {
	if t == nil {
		t = &CreateGroup_CreateGroup_Group{}
	}
	return &t.Region
}

type CreateGroup_CreateGroup struct {
	Group CreateGroup_CreateGroup_Group "json:\"group\" graphql:\"group\""
}

func (t *CreateGroup_CreateGroup) GetGroup() *CreateGroup_CreateGroup_Group {
	if t == nil {
		t = &CreateGroup_CreateGroup{}
	}
	return &t.Group
}

type DeleteGroup_DeleteGroup struct {
	DeletedID string "json:\"deletedID\" graphql:\"deletedID\""
}

func (t *DeleteGroup_DeleteGroup) GetDeletedID() string {
	if t == nil {
		t = &DeleteGroup_DeleteGroup{}
	}
	return t.DeletedID
}

type CreateDatabase struct {
	CreateDatabase CreateDatabase_CreateDatabase "json:\"createDatabase\" graphql:\"createDatabase\""
}

func (t *CreateDatabase) GetCreateDatabase() *CreateDatabase_CreateDatabase {
	if t == nil {
		t = &CreateDatabase{}
	}
	return &t.CreateDatabase
}

type DeleteDatabase struct {
	DeleteDatabase DeleteDatabase_DeleteDatabase "json:\"deleteDatabase\" graphql:\"deleteDatabase\""
}

func (t *DeleteDatabase) GetDeleteDatabase() *DeleteDatabase_DeleteDatabase {
	if t == nil {
		t = &DeleteDatabase{}
	}
	return &t.DeleteDatabase
}

type GetAllDatabases struct {
	Databases GetAllDatabases_Databases "json:\"databases\" graphql:\"databases\""
}

func (t *GetAllDatabases) GetDatabases() *GetAllDatabases_Databases {
	if t == nil {
		t = &GetAllDatabases{}
	}
	return &t.Databases
}

type GetDatabase struct {
	Database GetDatabase_Database "json:\"database\" graphql:\"database\""
}

func (t *GetDatabase) GetDatabase() *GetDatabase_Database {
	if t == nil {
		t = &GetDatabase{}
	}
	return &t.Database
}

type GetGroup struct {
	Group GetGroup_Group "json:\"group\" graphql:\"group\""
}

func (t *GetGroup) GetGroup() *GetGroup_Group {
	if t == nil {
		t = &GetGroup{}
	}
	return &t.Group
}

type GetAllGroups struct {
	Groups GetAllGroups_Groups "json:\"groups\" graphql:\"groups\""
}

func (t *GetAllGroups) GetGroups() *GetAllGroups_Groups {
	if t == nil {
		t = &GetAllGroups{}
	}
	return &t.Groups
}

type CreateGroup struct {
	CreateGroup CreateGroup_CreateGroup "json:\"createGroup\" graphql:\"createGroup\""
}

func (t *CreateGroup) GetCreateGroup() *CreateGroup_CreateGroup {
	if t == nil {
		t = &CreateGroup{}
	}
	return &t.CreateGroup
}

type DeleteGroup struct {
	DeleteGroup DeleteGroup_DeleteGroup "json:\"deleteGroup\" graphql:\"deleteGroup\""
}

func (t *DeleteGroup) GetDeleteGroup() *DeleteGroup_DeleteGroup {
	if t == nil {
		t = &DeleteGroup{}
	}
	return &t.DeleteGroup
}

const CreateDatabaseDocument = `mutation CreateDatabase ($input: CreateDatabaseInput!) {
	createDatabase(input: $input) {
		database {
			id
			name
			organizationID
			provider
			status
			dsn
			geo
		}
	}
}
`

func (c *Client) CreateDatabase(ctx context.Context, input CreateDatabaseInput, interceptors ...clientv2.RequestInterceptor) (*CreateDatabase, error) {
	vars := map[string]any{
		"input": input,
	}

	var res CreateDatabase
	if err := c.Client.Post(ctx, "CreateDatabase", CreateDatabaseDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const DeleteDatabaseDocument = `mutation DeleteDatabase ($name: String!) {
	deleteDatabase(name: $name) {
		deletedID
	}
}
`

func (c *Client) DeleteDatabase(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*DeleteDatabase, error) {
	vars := map[string]any{
		"name": name,
	}

	var res DeleteDatabase
	if err := c.Client.Post(ctx, "DeleteDatabase", DeleteDatabaseDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const GetAllDatabasesDocument = `query GetAllDatabases {
	databases {
		edges {
			node {
				id
				name
				organizationID
				provider
				status
				dsn
				geo
			}
		}
	}
}
`

func (c *Client) GetAllDatabases(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllDatabases, error) {
	vars := map[string]any{}

	var res GetAllDatabases
	if err := c.Client.Post(ctx, "GetAllDatabases", GetAllDatabasesDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const GetDatabaseDocument = `query GetDatabase ($name: String!) {
	database(name: $name) {
		id
		name
		organizationID
		provider
		status
		dsn
		geo
	}
}
`

func (c *Client) GetDatabase(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*GetDatabase, error) {
	vars := map[string]any{
		"name": name,
	}

	var res GetDatabase
	if err := c.Client.Post(ctx, "GetDatabase", GetDatabaseDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const GetGroupDocument = `query GetGroup ($name: String!) {
	group(name: $name) {
		id
		name
		description
		primaryLocation
		locations
		region
	}
}
`

func (c *Client) GetGroup(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*GetGroup, error) {
	vars := map[string]any{
		"name": name,
	}

	var res GetGroup
	if err := c.Client.Post(ctx, "GetGroup", GetGroupDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const GetAllGroupsDocument = `query GetAllGroups {
	groups {
		edges {
			node {
				id
				name
				description
				primaryLocation
				locations
				region
			}
		}
	}
}
`

func (c *Client) GetAllGroups(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetAllGroups, error) {
	vars := map[string]any{}

	var res GetAllGroups
	if err := c.Client.Post(ctx, "GetAllGroups", GetAllGroupsDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const CreateGroupDocument = `mutation CreateGroup ($input: CreateGroupInput!) {
	createGroup(input: $input) {
		group {
			id
			name
			description
			primaryLocation
			locations
			region
		}
	}
}
`

func (c *Client) CreateGroup(ctx context.Context, input CreateGroupInput, interceptors ...clientv2.RequestInterceptor) (*CreateGroup, error) {
	vars := map[string]any{
		"input": input,
	}

	var res CreateGroup
	if err := c.Client.Post(ctx, "CreateGroup", CreateGroupDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const DeleteGroupDocument = `mutation DeleteGroup ($name: String!) {
	deleteGroup(name: $name) {
		deletedID
	}
}
`

func (c *Client) DeleteGroup(ctx context.Context, name string, interceptors ...clientv2.RequestInterceptor) (*DeleteGroup, error) {
	vars := map[string]any{
		"name": name,
	}

	var res DeleteGroup
	if err := c.Client.Post(ctx, "DeleteGroup", DeleteGroupDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	CreateDatabaseDocument:  "CreateDatabase",
	DeleteDatabaseDocument:  "DeleteDatabase",
	GetAllDatabasesDocument: "GetAllDatabases",
	GetDatabaseDocument:     "GetDatabase",
	GetGroupDocument:        "GetGroup",
	GetAllGroupsDocument:    "GetAllGroups",
	CreateGroupDocument:     "CreateGroup",
	DeleteGroupDocument:     "DeleteGroup",
}
