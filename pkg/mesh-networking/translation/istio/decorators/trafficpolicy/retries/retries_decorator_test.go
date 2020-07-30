package retries_test

import (
	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/decorators/trafficpolicy"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/decorators/trafficpolicy/retries"
	"istio.io/api/networking/v1alpha3"
)

var _ = Describe("RetriesDecorator", func() {
	var (
		retriesDecorator trafficpolicy.VirtualServiceDecorator
		output           *v1alpha3.HTTPRoute
	)

	BeforeEach(func() {
		retriesDecorator = retries.NewRetriesDecorator()
		output = &v1alpha3.HTTPRoute{}
	})

	It("should set retries", func() {
		registerField := func(fieldPtr, val interface{}) error {
			return nil
		}
		appliedPolicy := &discoveryv1alpha2.MeshServiceStatus_AppliedTrafficPolicy{
			Spec: &v1alpha2.TrafficPolicySpec{
				Retries: &v1alpha2.TrafficPolicySpec_RetryPolicy{
					Attempts:      5,
					PerTryTimeout: &types.Duration{Seconds: 2},
				},
			},
		}
		expectedRetries := &v1alpha3.HTTPRetry{
			Attempts:      5,
			PerTryTimeout: &types.Duration{Seconds: 2},
		}
		err := retriesDecorator.ApplyToVirtualService(
			appliedPolicy,
			nil,
			output,
			registerField,
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(output.Retries).To(Equal(expectedRetries))
	})

	It("should not set retries if error during field registration", func() {
		testErr := eris.New("registration error")
		registerField := func(fieldPtr, val interface{}) error {
			return testErr
		}
		appliedPolicy := &discoveryv1alpha2.MeshServiceStatus_AppliedTrafficPolicy{
			Spec: &v1alpha2.TrafficPolicySpec{
				Retries: &v1alpha2.TrafficPolicySpec_RetryPolicy{
					Attempts:      5,
					PerTryTimeout: &types.Duration{Seconds: 2},
				},
			},
		}
		err := retriesDecorator.ApplyToVirtualService(
			appliedPolicy,
			nil,
			output,
			registerField,
		)
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})
})