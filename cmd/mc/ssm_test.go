package mc

//type mockSSMAWSClient struct {
//	startSessionResp *ssm.StartSessionOutput
//}
//
//func (m mockSSMAWSClient) StartSession(ctx context.Context, params *ssm.StartSessionInput, optFns ...func(options *ssm.Options)) (*ssm.StartSessionOutput, error) {
//	return m.startSessionResp, nil
//}
//
//func NewFakeK8sClient(t *testing.T, objs ...client.Object) client.Client {
//	s := runtime.NewScheme()
//	if err := corev1.AddToScheme(s); err != nil {
//		t.Fatal(err)
//	}
//
//	return fake.NewClientBuilder().WithScheme(s).WithObjects(objs...).Build()
//}
//
//func Test_SSMRun(t *testing.T) {
//	const (
//		fakeNodeName      = "ip-10-0-0-0.us-east-2.compute.internal"
//		fakeEC2InstanceID = "i-0000aaaabbbb"
//	)
//
//	tests := []struct {
//		name          string
//		node          string
//		ec2InstanceID string
//		objs          []client.Object
//		expectErr     bool
//	}{
//		{
//			name:          "happy path",
//			node:          fakeNodeName,
//			ec2InstanceID: fakeEC2InstanceID,
//			objs: []client.Object{
//				&corev1.Node{
//					ObjectMeta: metav1.ObjectMeta{
//						Name: fakeNodeName,
//					},
//					Spec: corev1.NodeSpec{
//						ProviderID: fmt.Sprintf("aws:///us-east-2a/%s", fakeEC2InstanceID),
//					},
//				},
//			},
//			expectErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			s := &ssmClient{
//				k8sClient: NewFakeK8sClient(t, test.objs...),
//				node:      fakeNodeName,
//				awsClient: mockSSMAWSClient{startSessionResp: &ssm.StartSessionOutput{}},
//			}
//
//			if err := s.Run(context.Background()); err != nil {
//				if !test.expectErr {
//					t.Errorf("expected no err, got %v", err)
//				}
//			} else {
//				if test.expectErr {
//					t.Error("expected err, got nil")
//				}
//			}
//		})
//	}
//}
