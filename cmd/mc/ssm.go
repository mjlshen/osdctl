package mc

//type ssmClient struct {
//	k8sClient     client.Client
//	clusterID     string
//	node          string
//	ec2InstanceID string
//	awsClient     awsClient
//}
//
//type awsClient interface {
//	StartSession(ctx context.Context, params *ssm.StartSessionInput, optFns ...func(options *ssm.Options)) (*ssm.StartSessionOutput, error)
//}
//
//func newCmdSSM() *cobra.Command {
//	s := &ssmClient{}
//
//	ssmCmd := &cobra.Command{
//		Use:          "ssm",
//		Short:        "Use AWS SSM to connect to an EC2 instance",
//		Long:         "Use AWS SSM to connect to an EC2 instance",
//		Example:      "osdctl mc ssm -c ${MC_CLUSTER_ID} ${NODE}",
//		SilenceUsage: true,
//		Args:         cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			if err := s.init(args); err != nil {
//				return err
//			}
//			return s.Run(context.Background())
//		},
//	}
//
//	ssmCmd.Flags().StringVarP(&s.clusterID, "cluster-id", "c", "", "OCM internal/external cluster id or cluster name of the ROSA HCP management cluster")
//
//	return ssmCmd
//}
//
//func (s *ssmClient) Run(ctx context.Context) error {
//	if s.ec2InstanceID == "" {
//		node := &corev1.Node{}
//		if err := s.k8sClient.Get(ctx, client.ObjectKey{Name: s.node}, node); err != nil {
//			return fmt.Errorf("failed to find node: %s on cluster: %v", s.node, err)
//		}
//
//		// The providerID is in the form aws:///%s/%s where the first string is the AWS AZ and the second
//		// string is the EC2 instance ID
//		s.ec2InstanceID = node.Spec.ProviderID[strings.LastIndex(node.Spec.ProviderID, "/")+1:]
//	}
//	log.Printf("Attempting to connect to: %s", s.ec2InstanceID)
//
//	resp, err := s.awsClient.StartSession(ctx, &ssm.StartSessionInput{
//		Target: aws.String(s.ec2InstanceID),
//		Reason: aws.String(fmt.Sprintf("Red Hat SRE via osdctl: %s", utils.Version)),
//	})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *ssmClient) init(args []string) error {
//	ocmClient, err := utils.CreateConnection()
//	if err != nil {
//		return err
//	}
//	defer ocmClient.Close()
//
//	cluster, err := utils.GetClusterAnyStatus(ocmClient, s.clusterID)
//	if err != nil {
//		return fmt.Errorf("failed to get OCM cluster info for %s: %s", s.clusterID, err)
//	}
//
//	cfg, err := osdCloud.CreateAWSV2Config(cluster)
//	if err != nil {
//		return err
//	}
//
//	scheme := runtime.NewScheme()
//	if err := corev1.AddToScheme(scheme); err != nil {
//		return err
//	}
//
//	k8sClient, err := k8s.New(cluster.ID(), client.Options{Scheme: scheme})
//	if err != nil {
//		return err
//	}
//
//	s.k8sClient = k8sClient
//	s.awsClient = ssm.NewFromConfig(cfg)
//	s.node = args[0]
//	log.Printf("%v", s)
//	return nil
//}
