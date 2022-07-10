<template>
  <div>
<button @click="GoBack()">Go back</button>
<button data-target="#createMessage" data-toggle="modal">Create message</button>
<div>Neprocitane poruke:</div>
<div  v-for="(message, index) in messages" :key="index" >
<button data-target="#showAndAnswer" data-toggle="modal" @click="FindMessages(message.FirstUsername)">
    {{message.FirstUsername}}
</button>
</div>

    <div class="modal fade" id="showAndAnswer" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-body" style="padding: 15px 50px">
                    <form role="form">

                            <div v-for="(message, index) in messages1" :key="index">
                              <h4 style="width: 600px" class="text">
                                {{message.FirstUsername}}: {{message.MessageText}}
                              </h4>
                            </div>
                        
                        <div class="form-group">
                            <label><span class="glyphicon glyphicon-eye-open"></span>Text:</label>
                            <input type="text" v-model="this.textMessage"/>
                        </div>
                        <button type="submit" class="btn btn-success btn-block" @click="Reply(textMessage)">
                            <span></span> Send Message
                        </button>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" data-dismiss="modal">Cancel</button>
                </div>
            </div>
        </div>
    </div>




    <div class="modal fade" id="createMessage" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-body" style="padding: 15px 50px">
                    <form role="form">
                        <select class="form-control" name="template" v-model="this.usernameForWho">
                      <option v-for="(conn, index) in connections" :key="index"  >
                            {{conn.SecondUsername}}
                      </option>
                  </select>
                        
                        <div class="form-group">
                            <label><span class="glyphicon glyphicon-eye-open"></span>Text:</label>
                            <input type="text" v-model="this.textMessage"/>
                        </div>
                        <button type="button" class="btn btn-success btn-block" @click="FindBlock(usernameForWho,textMessage)">
                            <span></span> Send Message
                        </button>
                        
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" data-dismiss="modal">Cancel</button>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>

import axios from 'axios';
export default {
  name: "Messages",
  data() {
    return {
        textMessage:"",
        usernameWhoSent:"",
        usernameForWho:"",
        connections:"",
        conn:{Id:0,FirstUsername:"",SecondUsername:"", IsConfirmed:false, IsDeleted:false},
        message:{Id:0,FirstUsername:"",SecondUsername:"",MessageText:"",DateCreated:"",IsRead:null},
        messages:"",
        Username:"",
        username1:"",
        messages1:"",
        messager:"",
        messages2:[],
              block:{
        Id:0,
        FirstUsername: "",
        SecondUsername: "",
        mess:""
      },
    };
  },


  methods: {
    async GoBack(){
         this.$router.push({ name: "StartPageUser" });
    },
    async FindMessages(FirstUsername) {
        this.Username = localStorage.getItem("username");
        this.username1=FirstUsername
        axios.get("http://localhost:9090/getMessages/"+this.Username+"/"+this.username1)
        .then (response => { 
        this.messages1 = response.data.message;
    })
    this.messages2.push(this.messages1);
    console.log("USPELA i duzina poruka je"+ this.messages2.length)

        axios.post("http://localhost:9090/messager",{
            message: this.messages2,
        })
    },
    async Reply(textMessage) {
        this.textMessage=textMessage;
        axios.post("http://localhost:9090/message",{
            FirstUsername: this.Username,
            SecondUsername: this.username1,
            MessageText: this.textMessage
        })
    },
        async FindBlock1(textMessage){
        this.textMessage=textMessage;
    axios.get("http://localhost:9090/blocked/"+this.Username+"/"+this.username1)
      .then (response => { 
        console.log("FIrstame je:"+response.data.block.FirstUsername)
        console.log("SecondUsername je:"+response.data.block.SecondUsername)
          this.block.FirstUsername = response.data.block.FirstUsername;
          this.block.SecondUsername = response.data.block.SecondUsername;
          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
              axios.get("http://localhost:9090/blocked/"+this.username1+"/"+this.Username)
                    .then (response => { 
                                this.block.FirstUsername = response.data.block.FirstUsername;
                                this.block.SecondUsername = response.data.block.SecondUsername;
                                          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
                                            this.Reply(this.textMessage);
                                            }
                    })
          } else {
            console.log("HERE AM I")
           this.mess="You can not send a message to this person!"
          }
      })    

    },

    async GetUsernamesFromConn(){
   this.usernameWhoSent = localStorage.getItem("username");
    axios.get("http://localhost:9090/getConnUsername/"+this.usernameWhoSent)
    .then (response => { 
    this.connections = response.data.connection;
    })
    },
        async CheckMessages() {
      this.Username = localStorage.getItem("username"); 
      axios.get("http://localhost:9090/getUnreadMessages/"+this.Username)
      .then (response => { 
      this.messages = response.data.message;
      console.log("Broj poruka je "+this.messages.length)
    })
    },
    async CreateMessage(usernameForWho,textMessage){
        this.usernameForWho=usernameForWho;
        this.textMessage=textMessage;
        console.log("KO SALJE"+this.usernameWhoSent)
         console.log("KO PRIMA"+this.usernameForWho)
         console.log("tekst"+this.textMessage)
        axios.post("http://localhost:9090/message",{
            FirstUsername: this.usernameWhoSent,
            SecondUsername: this.usernameForWho,
            MessageText: this.textMessage
        })
    },

        async FindBlock(usernameForWho,textMessage){
        this.usernameForWho=usernameForWho;
        this.textMessage=textMessage;
    axios.get("http://localhost:9090/blocked/"+this.usernameWhoSent+"/"+this.usernameForWho)
      .then (response => { 
        console.log("FIrstame je:"+response.data.block.FirstUsername)
        console.log("SecondUsername je:"+response.data.block.SecondUsername)
          this.block.FirstUsername = response.data.block.FirstUsername;
          this.block.SecondUsername = response.data.block.SecondUsername;
          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
              axios.get("http://localhost:9090/blocked/"+this.usernameForWho+"/"+this.usernameWhoSent)
                    .then (response => { 
                                this.block.FirstUsername = response.data.block.FirstUsername;
                                this.block.SecondUsername = response.data.block.SecondUsername;
                                          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
                                            this.CreateMessage(this.usernameForWho,this.textMessage);
                                            }
                    })
          } else {
            console.log("HERE AM I")
           this.mess="You can not send a message to this person!"
          }
      })    

    },
  },
  async created(){
    await this.GetUsernamesFromConn();
    await this.CheckMessages();
  }
};
</script>


<style scoped>
.row-boats {
  display: flex;
}
.text {
  text-align: left;
  font-size: 15px;
  font-weight: 600;
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-info {
  margin-left: 30px;
  margin-top: 10px;
}
.row {
  width: 660%;
  padding-left: 60px;
  height: 10%;
  margin-top: 100px;
  margin-left: 5000px;
  margin-right: 500px;
  opacity: 1.2;
  border-radius: 10px;
  align-content: center;
  background-color: rgba(255, 255, 255, 0.288);
  box-shadow: 0 5px 4px rgb(0 0 0 / 30%), 0 1px 1px rgb(0 0 0 / 22%);
}
/* kartica u okviru stavke menija  */
.containerInfo {
  width: 97%;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 10px;
  opacity: 1;
  border-radius: 10px;
  align-content: center;
  position: relative;
  background-color: fff;
  box-shadow: 0 19px 40px rgb(0 0 0 / 30%), 0 15px 12px rgb(0 0 0 / 22%);
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
</style>