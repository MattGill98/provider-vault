/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	secretbackendconnection "github.com/MattGill98/provider-vault/internal/controller/database/secretbackendconnection"
	secretbackendrole "github.com/MattGill98/provider-vault/internal/controller/database/secretbackendrole"
	providerconfig "github.com/MattGill98/provider-vault/internal/controller/providerconfig"
	mount "github.com/MattGill98/provider-vault/internal/controller/vault/mount"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secretbackendconnection.Setup,
		secretbackendrole.Setup,
		providerconfig.Setup,
		mount.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
