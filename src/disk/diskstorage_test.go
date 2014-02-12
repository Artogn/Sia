package disk

import "testing"
import "os"
func Test_Swarm(t *testing.T){
	i,err:=CreateSwarmSystem("SW1")
	defer os.Remove("SW1")
	if err!=nil{
		t.Error(err.Error())
	}
	if i==nil{
		t.Fatal("first returned object is nil")
	}

	b,err:=CreateSwarmSystem("SW2")
	defer os.Remove("SW2")
	if err!=nil{
		t.Error("Failed to create second swarm")
	}
	err=i.CreateFile("hi_there",1000)
	if err!=nil{
		t.Error(err.Error())
		t.Fatal("Failed to Create file")
		
	}
	b.CreateFile("Hello",1000)
	f1:=[]byte{12,34,51,23,51,12,51}
	f2:=[]byte{24,34,51,25}
	err=i.WriteFile("hi_there",10,f1)
	if nil!=err{
		if os.IsPermission(err){
			t.Error("Permission bits are set so that things don't work.")
		}else{
			t.Error("Failed to write file 1")
		}
	}
	
	if nil!=b.WriteFile("Hello",0,f2){
		t.Error("Error. File 2 failed to write.")
	}
	if nil!=i.DeleteFile("hi_there"){
		t.Error("Failed to destroy file")
	}
	if nil==b.DeleteFile("hello"){
		t.Error("Swarm did not cause error when incorrect name applied.")
	}

	
}
