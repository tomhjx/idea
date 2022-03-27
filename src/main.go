package main

import (
	"log"
	"time"
)

type ScoreBoard struct {
	duration time.Duration
}

func (i *ScoreBoard) Add(reporter *ScoreBoard) {
	i.duration += reporter.duration
}

func (i *ScoreBoard) Duration() time.Duration {
	return i.duration
}

type HttpApi struct {
	scoreBoard *ScoreBoard
}

func NewHttpApi() *HttpApi {
	return &HttpApi{scoreBoard: &ScoreBoard{}}
}

func (i *HttpApi) Response() *ScoreBoard {
	lb := NewHttpApiLoadBlancer()
	scoreBoard := lb.ProxyFastCGI(NewPHPFPMServer())
	i.scoreBoard.Add(scoreBoard)
	return i.scoreBoard
}

type Resources struct {
	CPUs   int //m, 1c = 1000m
	Memory int //byte
}

func (i *Resources) Add(r *Resources) {
	i.CPUs += r.CPUs
	i.Memory += r.Memory
}

type FastCGIContent struct {
}

type PHPProcessor struct {
	scoreBoard *ScoreBoard
	resources  *Resources
}

func NewPHPProcessor() *PHPProcessor {
	return &PHPProcessor{
		scoreBoard: &ScoreBoard{},
		resources:  &Resources{},
	}
}

func (i *PHPProcessor) ZendCompile() {
	log.Println("PHP Processor Compile By Zend.")
	i.scoreBoard.duration += 100 * time.Millisecond
	i.resources.CPUs += 10
	i.resources.Memory += 10
}

func (i *PHPProcessor) ZendExecute() {
	log.Println("PHP Processor Execute By Zend.")
	i.scoreBoard.duration += 100 * time.Millisecond
	i.resources.CPUs += 10
	i.resources.Memory += 32 * 1024 * 1024
}

func (i *PHPProcessor) Run() (*ScoreBoard, *Resources) {
	i.ZendCompile()
	i.ZendExecute()
	return i.scoreBoard, i.resources
}

type PHPFPMServer struct {
	scoreBoard *ScoreBoard
	logic      *PHPProcessor
}

func NewPHPFPMServer() *PHPFPMServer {
	return &PHPFPMServer{
		scoreBoard: &ScoreBoard{},
		logic:      NewPHPProcessor()}
}

func (i *PHPFPMServer) Accept() {

}

func (i *PHPFPMServer) Run() {
	s, _ := i.logic.Run()
	i.scoreBoard.Add(s)
}

func (i *PHPFPMServer) Response(content *FastCGIContent) *ScoreBoard {
	i.Run()
	return i.scoreBoard
}

type HttpApiLoadBlancer struct {
	scoreBoard *ScoreBoard
}

func NewHttpApiLoadBlancer() *HttpApiLoadBlancer {
	return &HttpApiLoadBlancer{scoreBoard: &ScoreBoard{}}
}

func (i *HttpApiLoadBlancer) Location() {
}

func (i *HttpApiLoadBlancer) ProxyFastCGI(backend *PHPFPMServer) *ScoreBoard {
	content := &FastCGIContent{}
	i.scoreBoard.Add(backend.Response(content))
	return i.scoreBoard
}

type HttpServer interface {
	TLS() string // 1.1,1.2,1.3
}

type ChromeBrowser struct {
}

func (i *ChromeBrowser) CallHttpApi(api *HttpApi) {
	r := api.Response()
	log.Printf("cost time: %d ms", r.Duration().Milliseconds())
}

// interaction{actor1, actor2}

func requestHttpsApi() {
	browser := &ChromeBrowser{}
	browser.CallHttpApi(NewHttpApi())
}

func main() {
	requestHttpsApi()

}
