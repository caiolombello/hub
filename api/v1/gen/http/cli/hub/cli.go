// Code generated by goa v3.2.2, DO NOT EDIT.
//
// hub HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	resourcec "github.com/tektoncd/hub/api/v1/gen/http/resource/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `resource (query|list|versions-by-id|by-catalog-kind-name-version|by-version-id|by-catalog-kind-name|by-id)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` resource query --name "buildah" --kinds '[
      "task",
      "pipelines"
   ]' --tags '[
      "image",
      "build"
   ]' --limit 100 --match "exact"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		resourceFlags = flag.NewFlagSet("resource", flag.ContinueOnError)

		resourceQueryFlags     = flag.NewFlagSet("query", flag.ExitOnError)
		resourceQueryNameFlag  = resourceQueryFlags.String("name", "", "")
		resourceQueryKindsFlag = resourceQueryFlags.String("kinds", "", "")
		resourceQueryTagsFlag  = resourceQueryFlags.String("tags", "", "")
		resourceQueryLimitFlag = resourceQueryFlags.String("limit", "1000", "")
		resourceQueryMatchFlag = resourceQueryFlags.String("match", "contains", "")

		resourceListFlags     = flag.NewFlagSet("list", flag.ExitOnError)
		resourceListLimitFlag = resourceListFlags.String("limit", "1000", "")

		resourceVersionsByIDFlags  = flag.NewFlagSet("versions-by-id", flag.ExitOnError)
		resourceVersionsByIDIDFlag = resourceVersionsByIDFlags.String("id", "REQUIRED", "ID of a resource")

		resourceByCatalogKindNameVersionFlags       = flag.NewFlagSet("by-catalog-kind-name-version", flag.ExitOnError)
		resourceByCatalogKindNameVersionCatalogFlag = resourceByCatalogKindNameVersionFlags.String("catalog", "REQUIRED", "name of catalog")
		resourceByCatalogKindNameVersionKindFlag    = resourceByCatalogKindNameVersionFlags.String("kind", "REQUIRED", "kind of resource")
		resourceByCatalogKindNameVersionNameFlag    = resourceByCatalogKindNameVersionFlags.String("name", "REQUIRED", "name of resource")
		resourceByCatalogKindNameVersionVersionFlag = resourceByCatalogKindNameVersionFlags.String("version", "REQUIRED", "version of resource")

		resourceByVersionIDFlags         = flag.NewFlagSet("by-version-id", flag.ExitOnError)
		resourceByVersionIDVersionIDFlag = resourceByVersionIDFlags.String("version-id", "REQUIRED", "Version ID of a resource's version")

		resourceByCatalogKindNameFlags       = flag.NewFlagSet("by-catalog-kind-name", flag.ExitOnError)
		resourceByCatalogKindNameCatalogFlag = resourceByCatalogKindNameFlags.String("catalog", "REQUIRED", "name of catalog")
		resourceByCatalogKindNameKindFlag    = resourceByCatalogKindNameFlags.String("kind", "REQUIRED", "kind of resource")
		resourceByCatalogKindNameNameFlag    = resourceByCatalogKindNameFlags.String("name", "REQUIRED", "Name of resource")

		resourceByIDFlags  = flag.NewFlagSet("by-id", flag.ExitOnError)
		resourceByIDIDFlag = resourceByIDFlags.String("id", "REQUIRED", "ID of a resource")
	)
	resourceFlags.Usage = resourceUsage
	resourceQueryFlags.Usage = resourceQueryUsage
	resourceListFlags.Usage = resourceListUsage
	resourceVersionsByIDFlags.Usage = resourceVersionsByIDUsage
	resourceByCatalogKindNameVersionFlags.Usage = resourceByCatalogKindNameVersionUsage
	resourceByVersionIDFlags.Usage = resourceByVersionIDUsage
	resourceByCatalogKindNameFlags.Usage = resourceByCatalogKindNameUsage
	resourceByIDFlags.Usage = resourceByIDUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "resource":
			svcf = resourceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "resource":
			switch epn {
			case "query":
				epf = resourceQueryFlags

			case "list":
				epf = resourceListFlags

			case "versions-by-id":
				epf = resourceVersionsByIDFlags

			case "by-catalog-kind-name-version":
				epf = resourceByCatalogKindNameVersionFlags

			case "by-version-id":
				epf = resourceByVersionIDFlags

			case "by-catalog-kind-name":
				epf = resourceByCatalogKindNameFlags

			case "by-id":
				epf = resourceByIDFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "resource":
			c := resourcec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "query":
				endpoint = c.Query()
				data, err = resourcec.BuildQueryPayload(*resourceQueryNameFlag, *resourceQueryKindsFlag, *resourceQueryTagsFlag, *resourceQueryLimitFlag, *resourceQueryMatchFlag)
			case "list":
				endpoint = c.List()
				data, err = resourcec.BuildListPayload(*resourceListLimitFlag)
			case "versions-by-id":
				endpoint = c.VersionsByID()
				data, err = resourcec.BuildVersionsByIDPayload(*resourceVersionsByIDIDFlag)
			case "by-catalog-kind-name-version":
				endpoint = c.ByCatalogKindNameVersion()
				data, err = resourcec.BuildByCatalogKindNameVersionPayload(*resourceByCatalogKindNameVersionCatalogFlag, *resourceByCatalogKindNameVersionKindFlag, *resourceByCatalogKindNameVersionNameFlag, *resourceByCatalogKindNameVersionVersionFlag)
			case "by-version-id":
				endpoint = c.ByVersionID()
				data, err = resourcec.BuildByVersionIDPayload(*resourceByVersionIDVersionIDFlag)
			case "by-catalog-kind-name":
				endpoint = c.ByCatalogKindName()
				data, err = resourcec.BuildByCatalogKindNamePayload(*resourceByCatalogKindNameCatalogFlag, *resourceByCatalogKindNameKindFlag, *resourceByCatalogKindNameNameFlag)
			case "by-id":
				endpoint = c.ByID()
				data, err = resourcec.BuildByIDPayload(*resourceByIDIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// resourceUsage displays the usage of the resource command and its subcommands.
func resourceUsage() {
	fmt.Fprintf(os.Stderr, `The resource service provides details about all kind of resources
Usage:
    %s [globalflags] resource COMMAND [flags]

COMMAND:
    query: Find resources by a combination of name, kind and tags
    list: List all resources sorted by rating and name
    versions-by-id: Find all versions of a resource by its id
    by-catalog-kind-name-version: Find resource using name of catalog & name, kind and version of resource
    by-version-id: Find a resource using its version's id
    by-catalog-kind-name: Find resources using name of catalog, resource name and kind of resource
    by-id: Find a resource using it's id

Additional help:
    %s resource COMMAND --help
`, os.Args[0], os.Args[0])
}
func resourceQueryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource query -name STRING -kinds JSON -tags JSON -limit UINT -match STRING

Find resources by a combination of name, kind and tags
    -name STRING: 
    -kinds JSON: 
    -tags JSON: 
    -limit UINT: 
    -match STRING: 

Example:
    `+os.Args[0]+` resource query --name "buildah" --kinds '[
      "task",
      "pipelines"
   ]' --tags '[
      "image",
      "build"
   ]' --limit 100 --match "exact"
`, os.Args[0])
}

func resourceListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource list -limit UINT

List all resources sorted by rating and name
    -limit UINT: 

Example:
    `+os.Args[0]+` resource list --limit 100
`, os.Args[0])
}

func resourceVersionsByIDUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource versions-by-id -id UINT

Find all versions of a resource by its id
    -id UINT: ID of a resource

Example:
    `+os.Args[0]+` resource versions-by-id --id 1
`, os.Args[0])
}

func resourceByCatalogKindNameVersionUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource by-catalog-kind-name-version -catalog STRING -kind STRING -name STRING -version STRING

Find resource using name of catalog & name, kind and version of resource
    -catalog STRING: name of catalog
    -kind STRING: kind of resource
    -name STRING: name of resource
    -version STRING: version of resource

Example:
    `+os.Args[0]+` resource by-catalog-kind-name-version --catalog "tektoncd" --kind "pipeline" --name "buildah" --version "0.1"
`, os.Args[0])
}

func resourceByVersionIDUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource by-version-id -version-id UINT

Find a resource using its version's id
    -version-id UINT: Version ID of a resource's version

Example:
    `+os.Args[0]+` resource by-version-id --version-id 1
`, os.Args[0])
}

func resourceByCatalogKindNameUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource by-catalog-kind-name -catalog STRING -kind STRING -name STRING

Find resources using name of catalog, resource name and kind of resource
    -catalog STRING: name of catalog
    -kind STRING: kind of resource
    -name STRING: Name of resource

Example:
    `+os.Args[0]+` resource by-catalog-kind-name --catalog "tektoncd" --kind "pipeline" --name "buildah"
`, os.Args[0])
}

func resourceByIDUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource by-id -id UINT

Find a resource using it's id
    -id UINT: ID of a resource

Example:
    `+os.Args[0]+` resource by-id --id 1
`, os.Args[0])
}