package services

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type AliACMSvs struct {
	Config constant.ClientConfig
	Env    string
}

func NewAliACMSvs(endpoint string, namespaceId string, accessKey string, secretKey string, env string) *AliACMSvs {
	var config constant.ClientConfig = constant.ClientConfig{
		Endpoint:    endpoint + ":8080", //获取nacos节点ip的服务地址
		NamespaceId: namespaceId,        //nacos命名空间
		AccessKey:   accessKey,
		SecretKey:   secretKey,

		TimeoutMs:      5 * 1000,  //http请求超时时间，单位毫秒
		ListenInterval: 30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
		//BeatInterval:   5 * 1000,  //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）

		CacheDir:             "./cache", //缓存目录
		LogDir:               "./log",   //日志目录
		UpdateThreadNum:      20,        //更新服务的线程数
		NotLoadCacheAtStart:  true,      //在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,      //当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	return &AliACMSvs{
		Config: config,
		Env:    env,
	}
}

func (this *AliACMSvs) Load(key string) interface{} {
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": this.Config,
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: key,
		Group:  this.Env,
	})

	fmt.Println("Get config：" + content)

	return content
}
