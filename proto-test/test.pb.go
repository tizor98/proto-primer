// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto-test/test.proto

package proto_test

import (
	proto_student "github.com/tizor98/proto-primer/proto-student"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TEST_STATUS int32

const (
	TEST_STATUS_PASS TEST_STATUS = 0
	TEST_STATUS_FAIL TEST_STATUS = 1
)

// Enum value maps for TEST_STATUS.
var (
	TEST_STATUS_name = map[int32]string{
		0: "PASS",
		1: "FAIL",
	}
	TEST_STATUS_value = map[string]int32{
		"PASS": 0,
		"FAIL": 1,
	}
)

func (x TEST_STATUS) Enum() *TEST_STATUS {
	p := new(TEST_STATUS)
	*p = x
	return p
}

func (x TEST_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TEST_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_test_test_proto_enumTypes[0].Descriptor()
}

func (TEST_STATUS) Type() protoreflect.EnumType {
	return &file_proto_test_test_proto_enumTypes[0]
}

func (x TEST_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TEST_STATUS.Descriptor instead.
func (TEST_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{0}
}

type ANSWER_STATUS int32

const (
	ANSWER_STATUS_CORRECT   ANSWER_STATUS = 0
	ANSWER_STATUS_INCORRECT ANSWER_STATUS = 1
)

// Enum value maps for ANSWER_STATUS.
var (
	ANSWER_STATUS_name = map[int32]string{
		0: "CORRECT",
		1: "INCORRECT",
	}
	ANSWER_STATUS_value = map[string]int32{
		"CORRECT":   0,
		"INCORRECT": 1,
	}
)

func (x ANSWER_STATUS) Enum() *ANSWER_STATUS {
	p := new(ANSWER_STATUS)
	*p = x
	return p
}

func (x ANSWER_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ANSWER_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_test_test_proto_enumTypes[1].Descriptor()
}

func (ANSWER_STATUS) Type() protoreflect.EnumType {
	return &file_proto_test_test_proto_enumTypes[1]
}

func (x ANSWER_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ANSWER_STATUS.Descriptor instead.
func (ANSWER_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{1}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetTestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetTestResponse) Reset() {
	*x = SetTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTestResponse) ProtoMessage() {}

func (x *SetTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTestResponse.ProtoReflect.Descriptor instead.
func (*SetTestResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{2}
}

func (x *SetTestResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetTestResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TestId   string `protobuf:"bytes,4,opt,name=testId,proto3" json:"testId,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{3}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type ClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ClientStreamResponse) Reset() {
	*x = ClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{4}
}

func (x *ClientStreamResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EnrollmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	TestId    string `protobuf:"bytes,2,opt,name=testId,proto3" json:"testId,omitempty"`
}

func (x *EnrollmentRequest) Reset() {
	*x = EnrollmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollmentRequest) ProtoMessage() {}

func (x *EnrollmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollmentRequest.ProtoReflect.Descriptor instead.
func (*EnrollmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{5}
}

func (x *EnrollmentRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *EnrollmentRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type GetStudentPerTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestId string `protobuf:"bytes,1,opt,name=testId,proto3" json:"testId,omitempty"`
}

func (x *GetStudentPerTestRequest) Reset() {
	*x = GetStudentPerTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentPerTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentPerTestRequest) ProtoMessage() {}

func (x *GetStudentPerTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentPerTestRequest.ProtoReflect.Descriptor instead.
func (*GetStudentPerTestRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{6}
}

func (x *GetStudentPerTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type TakeTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TakeTest:
	//
	//	*TakeTestRequest_TestId
	//	*TakeTestRequest_Answer
	TakeTest isTakeTestRequest_TakeTest `protobuf_oneof:"takeTest"`
}

func (x *TakeTestRequest) Reset() {
	*x = TakeTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTestRequest) ProtoMessage() {}

func (x *TakeTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTestRequest.ProtoReflect.Descriptor instead.
func (*TakeTestRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{7}
}

func (m *TakeTestRequest) GetTakeTest() isTakeTestRequest_TakeTest {
	if m != nil {
		return m.TakeTest
	}
	return nil
}

func (x *TakeTestRequest) GetTestId() string {
	if x, ok := x.GetTakeTest().(*TakeTestRequest_TestId); ok {
		return x.TestId
	}
	return ""
}

func (x *TakeTestRequest) GetAnswer() string {
	if x, ok := x.GetTakeTest().(*TakeTestRequest_Answer); ok {
		return x.Answer
	}
	return ""
}

type isTakeTestRequest_TakeTest interface {
	isTakeTestRequest_TakeTest()
}

type TakeTestRequest_TestId struct {
	TestId string `protobuf:"bytes,1,opt,name=testId,proto3,oneof"`
}

type TakeTestRequest_Answer struct {
	Answer string `protobuf:"bytes,4,opt,name=answer,proto3,oneof"`
}

func (*TakeTestRequest_TestId) isTakeTestRequest_TakeTest() {}

func (*TakeTestRequest_Answer) isTakeTestRequest_TakeTest() {}

type TakeTestResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevAnswer     ANSWER_STATUS `protobuf:"varint,42,opt,name=prev_answer,json=prevAnswer,proto3,enum=test.ANSWER_STATUS" json:"prev_answer,omitempty"`
	TotalPoints    int32         `protobuf:"varint,1,opt,name=totalPoints,proto3" json:"totalPoints,omitempty"`
	TotalQuestions int32         `protobuf:"varint,2,opt,name=totalQuestions,proto3" json:"totalQuestions,omitempty"`
	Status         TEST_STATUS   `protobuf:"varint,3,opt,name=status,proto3,enum=test.TEST_STATUS" json:"status,omitempty"`
}

func (x *TakeTestResults) Reset() {
	*x = TakeTestResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTestResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTestResults) ProtoMessage() {}

func (x *TakeTestResults) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTestResults.ProtoReflect.Descriptor instead.
func (*TakeTestResults) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{8}
}

func (x *TakeTestResults) GetPrevAnswer() ANSWER_STATUS {
	if x != nil {
		return x.PrevAnswer
	}
	return ANSWER_STATUS_CORRECT
}

func (x *TakeTestResults) GetTotalPoints() int32 {
	if x != nil {
		return x.TotalPoints
	}
	return 0
}

func (x *TakeTestResults) GetTotalQuestions() int32 {
	if x != nil {
		return x.TotalQuestions
	}
	return 0
}

func (x *TakeTestResults) GetStatus() TEST_STATUS {
	if x != nil {
		return x.Status
	}
	return TEST_STATUS_PASS
}

type TakeTestQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId      string        `protobuf:"bytes,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
	Question        string        `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	CurrentQuestion int32         `protobuf:"varint,3,opt,name=currentQuestion,proto3" json:"currentQuestion,omitempty"`
	TotalQuestions  int32         `protobuf:"varint,4,opt,name=totalQuestions,proto3" json:"totalQuestions,omitempty"`
	PrevAnswer      ANSWER_STATUS `protobuf:"varint,5,opt,name=prev_answer,json=prevAnswer,proto3,enum=test.ANSWER_STATUS" json:"prev_answer,omitempty"`
}

func (x *TakeTestQuestion) Reset() {
	*x = TakeTestQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTestQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTestQuestion) ProtoMessage() {}

func (x *TakeTestQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTestQuestion.ProtoReflect.Descriptor instead.
func (*TakeTestQuestion) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{9}
}

func (x *TakeTestQuestion) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *TakeTestQuestion) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *TakeTestQuestion) GetCurrentQuestion() int32 {
	if x != nil {
		return x.CurrentQuestion
	}
	return 0
}

func (x *TakeTestQuestion) GetTotalQuestions() int32 {
	if x != nil {
		return x.TotalQuestions
	}
	return 0
}

func (x *TakeTestQuestion) GetPrevAnswer() ANSWER_STATUS {
	if x != nil {
		return x.PrevAnswer
	}
	return ANSWER_STATUS_CORRECT
}

type TakeTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TakeTest:
	//
	//	*TakeTestResponse_Question
	//	*TakeTestResponse_Results
	TakeTest isTakeTestResponse_TakeTest `protobuf_oneof:"takeTest"`
}

func (x *TakeTestResponse) Reset() {
	*x = TakeTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTestResponse) ProtoMessage() {}

func (x *TakeTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTestResponse.ProtoReflect.Descriptor instead.
func (*TakeTestResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{10}
}

func (m *TakeTestResponse) GetTakeTest() isTakeTestResponse_TakeTest {
	if m != nil {
		return m.TakeTest
	}
	return nil
}

func (x *TakeTestResponse) GetQuestion() *TakeTestQuestion {
	if x, ok := x.GetTakeTest().(*TakeTestResponse_Question); ok {
		return x.Question
	}
	return nil
}

func (x *TakeTestResponse) GetResults() *TakeTestResults {
	if x, ok := x.GetTakeTest().(*TakeTestResponse_Results); ok {
		return x.Results
	}
	return nil
}

type isTakeTestResponse_TakeTest interface {
	isTakeTestResponse_TakeTest()
}

type TakeTestResponse_Question struct {
	Question *TakeTestQuestion `protobuf:"bytes,1,opt,name=question,proto3,oneof"`
}

type TakeTestResponse_Results struct {
	Results *TakeTestResults `protobuf:"bytes,2,opt,name=results,proto3,oneof"`
}

func (*TakeTestResponse_Question) isTakeTestResponse_TakeTest() {}

func (*TakeTestResponse_Results) isTakeTestResponse_TakeTest() {}

var File_proto_test_test_proto protoreflect.FileDescriptor

var file_proto_test_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x66, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22,
	0x49, 0x0a, 0x11, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x51,
	0x0a, 0x0f, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x22, 0xbc, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52,
	0x0a, 0x70, 0x72, 0x65, 0x76, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x45, 0x53,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xd6, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x0a, 0x70,
	0x72, 0x65, 0x76, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x54, 0x61,
	0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x6b,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x2a, 0x21, 0x0a, 0x0b, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x2a, 0x2b, 0x0a, 0x0d, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x52, 0x52, 0x45,
	0x43, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43,
	0x54, 0x10, 0x01, 0x32, 0xf8, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x47, 0x0a, 0x0e,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x08, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x7a,
	0x6f, 0x72, 0x39, 0x38, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_test_test_proto_rawDescOnce sync.Once
	file_proto_test_test_proto_rawDescData = file_proto_test_test_proto_rawDesc
)

func file_proto_test_test_proto_rawDescGZIP() []byte {
	file_proto_test_test_proto_rawDescOnce.Do(func() {
		file_proto_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_test_proto_rawDescData)
	})
	return file_proto_test_test_proto_rawDescData
}

var file_proto_test_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_test_test_proto_goTypes = []interface{}{
	(TEST_STATUS)(0),                 // 0: test.TEST_STATUS
	(ANSWER_STATUS)(0),               // 1: test.ANSWER_STATUS
	(*Test)(nil),                     // 2: test.Test
	(*GetTestRequest)(nil),           // 3: test.GetTestRequest
	(*SetTestResponse)(nil),          // 4: test.SetTestResponse
	(*Question)(nil),                 // 5: test.Question
	(*ClientStreamResponse)(nil),     // 6: test.ClientStreamResponse
	(*EnrollmentRequest)(nil),        // 7: test.EnrollmentRequest
	(*GetStudentPerTestRequest)(nil), // 8: test.GetStudentPerTestRequest
	(*TakeTestRequest)(nil),          // 9: test.TakeTestRequest
	(*TakeTestResults)(nil),          // 10: test.TakeTestResults
	(*TakeTestQuestion)(nil),         // 11: test.TakeTestQuestion
	(*TakeTestResponse)(nil),         // 12: test.TakeTestResponse
	(*proto_student.Student)(nil),    // 13: student.Student
}
var file_proto_test_test_proto_depIdxs = []int32{
	1,  // 0: test.TakeTestResults.prev_answer:type_name -> test.ANSWER_STATUS
	0,  // 1: test.TakeTestResults.status:type_name -> test.TEST_STATUS
	1,  // 2: test.TakeTestQuestion.prev_answer:type_name -> test.ANSWER_STATUS
	11, // 3: test.TakeTestResponse.question:type_name -> test.TakeTestQuestion
	10, // 4: test.TakeTestResponse.results:type_name -> test.TakeTestResults
	3,  // 5: test.TestService.GetTest:input_type -> test.GetTestRequest
	2,  // 6: test.TestService.SetTest:input_type -> test.Test
	5,  // 7: test.TestService.SetQuestions:input_type -> test.Question
	7,  // 8: test.TestService.EnrollStudents:input_type -> test.EnrollmentRequest
	8,  // 9: test.TestService.GetStudentsPerTest:input_type -> test.GetStudentPerTestRequest
	9,  // 10: test.TestService.TakeTest:input_type -> test.TakeTestRequest
	2,  // 11: test.TestService.GetTest:output_type -> test.Test
	4,  // 12: test.TestService.SetTest:output_type -> test.SetTestResponse
	6,  // 13: test.TestService.SetQuestions:output_type -> test.ClientStreamResponse
	6,  // 14: test.TestService.EnrollStudents:output_type -> test.ClientStreamResponse
	13, // 15: test.TestService.GetStudentsPerTest:output_type -> student.Student
	12, // 16: test.TestService.TakeTest:output_type -> test.TakeTestResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_test_test_proto_init() }
func file_proto_test_test_proto_init() {
	if File_proto_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentPerTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTestResults); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTestQuestion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_test_test_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*TakeTestRequest_TestId)(nil),
		(*TakeTestRequest_Answer)(nil),
	}
	file_proto_test_test_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*TakeTestResponse_Question)(nil),
		(*TakeTestResponse_Results)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_test_test_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_test_test_proto_goTypes,
		DependencyIndexes: file_proto_test_test_proto_depIdxs,
		EnumInfos:         file_proto_test_test_proto_enumTypes,
		MessageInfos:      file_proto_test_test_proto_msgTypes,
	}.Build()
	File_proto_test_test_proto = out.File
	file_proto_test_test_proto_rawDesc = nil
	file_proto_test_test_proto_goTypes = nil
	file_proto_test_test_proto_depIdxs = nil
}