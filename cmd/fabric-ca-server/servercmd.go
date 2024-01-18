package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerCmd Fabric CA 서버에 대한 명령줄 인터페이스와 Fabric CA 서버에서 사용하는 구성을 제공하는 cobra 명령을 캡슐화합니다.
type ServerCmd struct {
	// name 명령의 이름(init , start, version)
	name string
	// rootCmd 코브라 명령입니다
	rootCmd *cobra.Command
	// myViper 바이퍼 인스턴스
	myViper *viper.Viper
	// blockingStart 서버를 시작한 후 차단할지 여부를 나타냅니다.
	blockingStart bool
	// cfgFileName 구성 파일의 이름입니다.
	cfgFileName string
	// homeDirectory 서버의 홈 디렉토리 위치입니다.
	homeDirectory string
	// cfg 서버 구성입니다.
	//cfg *lib.ServerConfig
}

// NewCommand 실행 준비가 된 새 ServerCmd를 반환합니다
func NewCommand(name string, blockingStart bool) *ServerCmd {
	s := &ServerCmd{
		name:          name,
		blockingStart: blockingStart,
		myViper:       viper.New(),
	}
	s.init()
	return s
}

// init  ServerCmd 인스턴스를 초기화합니다. cobra 루트 및 하위 명령을 초기화하고 명령 flgs를 viper에 등록합니다.
func (s *ServerCmd) init() {
	rootCmd := &cobra.Command{
		Use:   cmdName,
		Short: longName,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := s.configInit()
			if err != nil {
				return err
			}
			cmd.SilenceUsage = true
			// util.CmdRunBegin(s.myViper)
			return nil
		},
	}
	s.rootCmd = rootCmd
}

// configRequired "version" 명령어는 설정파일이 필요하지 않음을 반환한다.
func (s *ServerCmd) configRequired() bool {
	return s.name != "version"
}
