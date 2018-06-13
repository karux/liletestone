package server

import (
	// "errors"
	"log"
	"context"
	"strconv"
	"strings"
	"fmt"
	"runtime"
	"github.com/karux/liletestone"

	opentracing "github.com/opentracing/opentracing-go"

)

func (s LiletestoneServer) Get(ctx context.Context, r *liletestone.GetRequest) (*liletestone.GetResponse, error) {
//	return nil, errors.New("not yet implemented")
  id := r.Id
	log.Print("request id ",id)
	retval := liletestone.GetResponse{}
	var counter int

	if (id != "") {
		// get the counter at the end of the string
		// convention   name ctr
		lastBin := strings.LastIndex( id, "-" )
		if (lastBin != -1) {
			strCounter := id[lastBin+1:]
			counter,_ = strconv.Atoi(strCounter)
			//fmt.Println("counter:",counter)
			counter++
			//fmt.Println("counter:",counter)
			id = id[0:lastBin]
		}

		next := strconv.Itoa(counter)
		id = id + "-" + next
	} else {
		retval.Id = "chrisTestWorld-"  + strconv.Itoa(counter)
	}

	if (counter < 10) {
		log.Println("current counter",counter)
		log.Println("upstream request: ",id)
		// call out to the upstream grpc server
		r.Id = id
		retval.Id = callGet(ctx, r.Id,counter)
	} else {
		retval.Id = id
	}

	log.Print("return response",retval.Id)
	return &retval, nil
}

func callGet(ctx context.Context, requestId string, counter int ) (responseId string) {
	defer func() {
			 if r := recover(); r != nil {
				 	 log.Println(identifyPanic())
					 log.Println("Recovered in f", r)
			 }
	 }()

	log.Print("calling upstream server in Get handler... with requestId "+requestId)
	var v liletestone.GetRequest = liletestone.GetRequest{}
	v.Id = requestId
	cli := liletestone.GetLiletestoneClient()

	// add some context to the Trace (if the Span is available)
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.SetTag("spCounter", counter)
	} else {
		log.Println("missing span")
	}
	span.SetTag("grpcMethod", "get")
	span.SetTag("upstream",true)

	template := opentracing.StartSpan("template.rendering", opentracing.ChildOf(span.Context()))
	defer template.Finish()
//	trace := NewClientTrace(span)

	context := opentracing.ContextWithSpan(ctx, template)

	resp, err := cli.Get(context, &v)
	if (err == nil) {
		if (resp != nil) {
			log.Println("responseId",resp.Id)
			responseId = resp.Id
		}
	} else {
		log.Print(err)
	}
	log.Print("returning ",responseId)
	return responseId
}


func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}
