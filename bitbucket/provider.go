package bitbucket

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gobb "github.com/ktrysmt/go-bitbucket"

	v1 "github.com/zahiar/terraform-provider-bitbucket/bitbucket/api/v1"
)

// Provider will create the necessary terraform provider to talk to the Bitbucket APIs you should
// specify a USERNAME and PASSWORD
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("BITBUCKET_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BITBUCKET_PASSWORD", nil),
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"bitbucket_hook":                resourceHook(),
			"bitbucket_default_reviewers":   resourceDefaultReviewers(),
			"bitbucket_repository":          resourceRepository(),
			"bitbucket_repository_variable": resourceRepositoryVariable(),
			"bitbucket_project":             resourceProject(),
			"bitbucket_branch_restriction":  resourceBranchRestriction(),
			"bitbucket_deployment":          resourceDeployment(),
			"bitbucket_deployment_variable": resourceDeploymentVariable(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bitbucket_user":       dataUser(),
			"bitbucket_workspace":  dataSourceBitbucketWorkspace(),
			"bitbucket_repository": dataSourceBitbucketRepository(),
		},
	}
}

type Clients struct {
	V1 *v1.Client
	V2 *gobb.Client
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := gobb.NewBasicAuth(
		d.Get("username").(string),
		d.Get("password").(string),
	)

	v1Client := v1.NewClient(
		&v1.Auth{
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
		},
	)

	clients := &Clients{
		V1: v1Client,
		V2: client,
	}

	return clients, nil
}
