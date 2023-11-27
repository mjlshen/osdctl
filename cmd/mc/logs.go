package mc

//
//import "github.com/spf13/cobra"
//
//type logs struct {
//	namespace string
//}
//
//func newCmdLogs() *cobra.Command {
//	l := &logs{}
//
//	logsCmd := &cobra.Command{
//		Use:     "logs",
//		Short:   "View ROSA HCP Management Cluster pod logs",
//		Long:    "View ROSA HCP Management Cluster pod logs",
//		Example: "osdctl mc logs -n ${NS} ${POD} -c ${CONTAINER}",
//		Args:    cobra.MaximumNArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			return l.Run(args)
//		},
//	}
//
//	logsCmd.Flags().StringVarP(l.namespace, "namespace", "n", "", "If present, the namespace scope to view logs for")
//
//	return logsCmd
//}
//
//func (l *logs) Run(args []string) error {
//	return nil
//}
