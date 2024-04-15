package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/pingidentity/terraform-provider-pingone/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New("dev")()
}
