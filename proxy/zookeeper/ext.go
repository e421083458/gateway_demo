package zookeeper

//
////注册服务
//func (z *ZkManager) RegistServer(module, host string) (err error) {
//	nodePath := z.pathPrefix + module
//	return z.RegistServerPath(nodePath, host)
//}
//
//func (z *ZkManager) GetServerList(module string) (list []string, err error) {
//	return z.GetServerListByPath(z.pathPrefix + module)
//}
//
//func (z *ZkManager) WatchServerList(module string) (chan []string, chan error) {
//	return z.WatchServerListByPath(z.pathPrefix + module)
//}
//
////watch机制，监听节点值变化
//func (z *ZkManager) WatchGetData(module string) (chan []byte, chan error) {
//	nodePath := z.pathPrefix + "config_" + module
//	return z.WatchPathData(nodePath)
//}
//
////获取配置
//func (z *ZkManager) GetData(module string) ([]byte, *zk.Stat, error) {
//	nodePath := z.pathPrefix + "config_" + module
//	return z.GetPathData(nodePath)
//}
//
////更新配置
//func (z *ZkManager) SetData(module string, config []byte, version int32) (err error) {
//	nodePath := z.pathPrefix + "config_" + module
//	return z.SetPathData(nodePath, config, version)
//}
