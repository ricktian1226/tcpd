package codec

import (
	//"fmt"
	"io"
	//"reflect"
	//"time"

	proto "code.google.com/p/goprotobuf/proto"
	"github.com/funny/binary"
	"github.com/funny/link"
)

func ProtoBuf(spliter binary.Spliter) link.CodecType {
	return protoBufCodecType{spliter}
}

type protoBufCodecType struct {
	Spliter binary.Spliter
}

func (codecType protoBufCodecType) NewEncoder(w io.Writer) link.Encoder {
	return protoBufEncoder{
		codecType.Spliter,
		binary.NewWriter(w),
	}
}

func (codecType protoBufCodecType) NewDecoder(r io.Reader) link.Decoder {
	return protoBufDecoder{
		codecType.Spliter,
		binary.NewReader(r),
	}
}

type protoBufEncoder struct {
	Spliter binary.Spliter
	Writer  *binary.Writer
}

func (encoder protoBufEncoder) Encode(msg interface{}) (err error) {

	//fmt.Printf("%v", reflect.TypeOf(msg))

	//time.Sleep(time.Second * 40)

	//转换成proto.Message类型
	var data []byte
	data, err = proto.Marshal(msg.(proto.Message))
	if err != nil {
		return err
	}

	encoder.Writer.WritePacket(data, encoder.Spliter)

	//fmt.Printf("protoBufEncoder.WritePacket : %v", data)

	return encoder.Writer.Flush()
}

type protoBufDecoder struct {
	Spliter binary.Spliter
	Reader  *binary.Reader
}

func (decoder protoBufDecoder) Decode(msg interface{}) (err error) {
	var data []byte
	//*(msg.(*[]byte)) = decoder.Reader.ReadPacket(decoder.Spliter)
	data = decoder.Reader.ReadPacket(decoder.Spliter)
	err = proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		return
	}

	return decoder.Reader.Error()
}
