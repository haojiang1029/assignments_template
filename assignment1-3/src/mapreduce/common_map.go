package mapreduce

import (
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

// doMap does the job of a map worker: it reads one of the input files
// (inFile), calls the user-defined map function (mapF) for that file's
// contents, and partitions the output into nReduce intermediate files.
func doMap(
	jobName string, // the name of the MapReduce job
	mapTaskNumber int, // which map task this is
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(file string, contents string) []KeyValue,
) {
	// TODO:
	// You will need to write this function.
	// You can find the filename for this map task's input to reduce task number
	// r using reduceName(jobName, mapTaskNumber, r). The ihash function (given
	// below doMap) should be used to decide which file a given key belongs into.
	//
	//var path = "C:\\Users\\hjian8\\go\\src\\mapreduce\\"
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}
	data := string(contents)
	res := mapF(inFile, data)

	encoders := make([]*json.Encoder, nReduce)
	for i:=0; i < nReduce; i++ {
		//tempFile := path + reduceName(jobName, mapTaskNumber, i)
		file, err := os.Create(reduceName(jobName, mapTaskNumber, i))
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		encoders[i] = json.NewEncoder(file)
		
	}
    for _, kv := range res {
    	hashKey := ihash(kv.Key) % uint32(nReduce)
    	err := encoders[hashKey].Encode(&kv)
		if err != nil {
			log.Fatal(err)
		}
		
	}
		

	// for _, kv := range res {
	// 	reduceTaskNumber := int(ihash(kv.Key)) % nReduce
	// 	tempFile := path + reduceName(jobName, mapTaskNumber, reduceTaskNumber)
	// 	file, err := os.Open(tempFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	enc := json.NewEncoder(file)
	// 	enc.Encode(kv)

	// }
	// The intermediate output of a map task is stored in the file
	// system as multiple files whose name indicates which map task produced
	// them, as well as which reduce task they are for. Coming up with a
	// scheme for how to store the key/value pairs on disk can be tricky,
	// especially when taking into account that both keys and values could
	// contain newlines, quotes, and any other character you can think of.
	//
	// One format often used for serializing data to a byte stream that the
	// other end can correctly reconstruct is JSON. You are not required to
	// use JSON, but as the output of the reduce tasks *must* be JSON,
	// familiarizing yourself with it here may prove useful. You can write
	// out a data structure as a JSON string to a file using the commented
	// code below. The corresponding decoding functions can be found in
	// common_reduce.go.
	//
	//   enc := json.NewEncoder(file)
	//   for _, kv := ... {
	//     err := enc.Encode(&kv)
	//
	// Remember to close the file after you have written all the values!
}

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
