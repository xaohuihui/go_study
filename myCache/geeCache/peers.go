package geeCache

// author: songyanhui
// datetime: 2022/1/20 15:34:47
// software: GoLand

import pb "geeCache/geecachepb"

/*
使用一致性hash选择节点        是                                     是
   |-------> 是否是远程节点 -----> HTTP 客户端访问远程节点 ---> 成功？ ---> 服务端返回返回值
				| 否										 |否
				|------------------------------------> 回退到本地节点处理.
 */

// PeerPicker 接口，PickPeer方法用于根据传入key选择相应节点(PeerGetter)
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 接口，Get方法用户从对应group查找缓存值， 对应上述流程中的HTTP客户端
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
