package main

import (
	"benc2proto/proto_structs"
	"fmt"
)

func main() {
	//k := Metainfo{}
	mi := proto_structs.MetaInfo{
		Info: &proto_structs.MetaInfo_Info{
			Name:        "name123",
			PieceLength: 1324,
			Pieces:      []string{"adf", "fsdfd"},
			Data: &proto_structs.MetaInfo_Info_Files{
				Files: &proto_structs.FileInfos{Infos: []*proto_structs.FileInfo{
					{Length: 999, Path: "path/to/file"},
				}},
			},
		},
		Announce: "announce",
	}
	fmt.Println(mi.String())
}
