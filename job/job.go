package job

import (
	"encoding/json"
	e "github.com/wercker/pkg/envvar"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Job represents a wercker job. See wercker/web/src/models/job.js for all
// available fields.
type Job struct {
	ID                   bson.ObjectId `bson:"_id" json:"_id"`
	MaxConcurrentJobs    int           `bson:"maxConcurrentJobs" json:"maxConcurrentJobs"`
	Payload              *Payload      `bson:"payload" json:"payload"`
	PipelineName         string        `bson:"pipelineName" json:"pipelineName"`
	PipelineType         string        `bson:"pipelineType" json:"pipelineType"`
	Priority             int           `bson:"priority" json:"priority"`
	ProjectID            string        `bson:"projectId" json:"projectId"`
	ProjectName          string        `bson:"projectName" json:"projectName"`
	ProjectOwnerID       string        `bson:"projectOwnerId" json:"projectOwnerId"`
	ProjectOwner         string        `bson:"projectOwner" json:"projectOwner"`
	RunID                bson.ObjectId `bson:"runId" json:"runId"`
	WerckerObjectID      string        `bson:"werckerObjectId" json:"werckerObjectId"`
	TargetID             bson.ObjectId `bson:"targetId" json:"targetId"`
	TargetName           string        `bson:"targetName" json:"targetName"`
	Version              int           `bson:"version" json:"version"`
	VppLabels            []string      `bson:"vppLabels" json:"vppLabels"`
	WorkflowCreationDate time.Time     `bson:"workflowCreationDate" json:"workflowCreationDate"`
	EnvVars              []*e.EnvVar   `bson:"envVars" json:"envVars"`
	Running              bool          `bson:"running" json:"running"`
	CreatedAt            time.Time     `bson:"createdAt" json:"createdAt"`
	SessionToken         string        `bson:"sessionToken" json:"sessionToken"`
}

// PayloadAppSettings represents a wercker job's payload application settings
type PayloadAppSettings struct {
	Env []*e.EnvVar `json:"env"`
}

// Payload represents a wercker job's payload
type Payload struct {
	ApplicationID            bson.ObjectId      `json:"applicationId"`
	ApplicationName          string             `json:"applicationName"`
	ApplicationSettings      PayloadAppSettings `json:"applicationSettings"`
	CacheURL                 string             `json:"cacheUrl"`
	GitURL                   string             `json:"gitUrl"`
	ApplicationOwnerName     string             `json:"applicationOwnerName"`
	Action                   string             `json:"action"`
	Branch                   string             `json:"branch"`
	Commit                   string             `json:"commit"`
	EnvVars                  []*e.EnvVar        `json:"envVars"`
	PipelineName             string             `json:"pipelineName"`
	PullRequestNumber        *int64             `json:"pullRequestNumber,omitempty"`
	RunID                    bson.ObjectId      `json:"runId"`
	StartedBy                string             `json:"startedBy"`
	TargetID                 bson.ObjectId      `json:"targetId"`
	TargetName               string             `json:"targetName"`
	WorkflowCreationDate     time.Time          `json:"workflowCreationDate"`
	SSHKeyPublic             string             `json:"sshKeyPublic"`
	SSHKeyPrivate            string             `json:"sshKey"`
	BuildWerckerYamlContents string             `json:"buildWerckerYamlContents,omitempty"`
	DeployTargetName         string             `json:"deployTargetName,omitempty"`
	PackageURL               string             `json:"packageUrl,omitempty"`
}

// GetBSON returns a string holding JSON representation of a Payload
func (p *Payload) GetBSON() (interface{}, error) {
	payloadString, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return string(payloadString), nil
}

// SetBSON decodes a string holding JSON representation of a Payload
func (p *Payload) SetBSON(raw bson.Raw) error {
	var s string
	raw.Unmarshal(&s)
	return json.Unmarshal([]byte(s), p)
}
