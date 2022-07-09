<template>
  <div>
    <div>   <button @click="Update()">Update profile info</button></div>
    <div>   <button @click="ResetPassword()">Reset Password</button></div>
    <div>   <button @click="CreatePost()">Create post</button></div>
    <div>   <button @click="Find()">Find public user</button></div>
    <div>   <button @click="JobOffer()">Create a job offer</button></div>
    <div>   <button @click="JobOffers()">See all job offers</button></div>
    <div>   <button @click="CreateConnection()">Search profile and follow</button></div>
    <div>   <button @click="Messages()">Messages</button></div>

    <div>   <button v-if="this.newUser.TypeOfProfile=='PRIVATE'" @click="SeeAllFollowRequests()">See all your friend requests</button></div>
  </div>
</template>


<script>
import axios from 'axios'
import Swal from 'sweetalert2';

export default {
  name: "StartPageUser",
   data: () => ({
    err: "" ,
    Username: "",
    messages:"",
    type:0,
          newUser: {
        Username: "",
        Password:"",
        Email:"",
        PhoneNumber:"",
        FirstName:"",
        LastName:"",
        DateOfBirth:"",
        Gender:"",
        TypeOfProfile:"",
        TypeOfUser:"",
        Biography:"",
        WorkExperience:"",
        Education:"",
        Skills:"",
        Interest:"",
        message:{Id:0,FirstUsername:"",SecondUsername:"",MessageText:"",DateCreated:"",IsRead:null}
      },
  }),
  methods: {
    async Messages(){
      this.$router.push({ name: "Messages" });
    },
    async getUser() {
               const headers ={
            Authorization: "Bearer " + this.token,
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json',
          }
  axios.get("http://localhost:9090/user/"+ this.Username, {
      headers}
    )
    .then((response) => {
      this.newUser.Username = response.data.user.Username
      this.newUser.Email = response.data.user.Email
      this.newUser.PhoneNumber = response.data.user.PhoneNumber
      this.newUser.FirstName = response.data.user.FirstName
      this.newUser.LastName = response.data.user.LastName
      this.newUser.Gender= response.data.user.Gender 
      this.newUser.TypeOfProfile = response.data.user.TypeOfProfile

      this.newUser.TypeOfUser = response.data.user.TypeOfUser
      this.newUser.Biography = response.data.user.TypeOfUser
      this.newUser.WorkExperience = response.data.user.WorkExperience
      this.newUser.Education = response.data.user.Education
      this.newUser.Skills = response.data.user.Skills
      this.newUser.Interest = response.data.user.Interest
        })

    },
    async CreateConnection() {
      this.$router.push({ name: "Connections" });
    },
    async CheckMessages() {
      this.Username = localStorage.getItem("username"); 
      axios.get("http://localhost:9090/getUnreadMessages/"+this.Username)
      .then (response => { 
      this.messages = response.data.message;
     if (this.messages.length>0) {
        console.log("USLA")
           new Swal({
             title:"Obavestenje!",
             type: "warning",
             text:'Imate novu poruku!',
           })  
      }
    })
    },
      async Update() {
          this.$router.push({ name: "UpdateInfoView" });
      },
      async CreatePost() {
          this.$router.push({ name: "CreatePost" });
      },
      async SeeAllFollowRequests() {
          this.$router.push({ name: "FollowRequests" });
      },
  async Find() {
    this.$router.push({ name: "FindPublicUserByLogUser" });
  },
async JobOffer() {
  this.$router.push({ name: "CreateJobOffer" });
},
async JobOffers() {
  this.$router.push({ name: "SeeJobOffers" });
}
  },
  async created() {
  this.Username = localStorage.getItem("username"); 
  this.token = localStorage.getItem("token")
  console.log("username je"+ this.Username);
  console.log("token je"+ this.token);
  this.type = localStorage.getItem("userType")
  console.log("type je"+ this.type);
   await this.getUser(this.Username);
   await this.CheckMessages();
  // console.log("TIP JE111 "+this.newUser.TypeOfUser);
  // localStorage.setItem("userType",this.newUser.TypeOfUser);
  // this.type = localStorage.getItem("userType");
  }
  
}
</script>

<style>
* {
  box-sizing: border-box;
  padding: 1px;
  font-family: Arial, Helvetica, sans-serif;
}

#heading1{
    text-align: center;
    padding: 30px;
}
img{
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
 
}
.myForm{
    max-width:500px;
   margin: auto;
   margin-top: 10px;
  }
  .class1 {
    
    display: flex;
    width: 100%;
    margin-bottom: 15px;
    
  }
  .icon {
    padding: 10px;
    background: darkseagreen;
    color:white;
    min-width: 50px;
    text-align: center;
  }
  .input-field {
    width: 100%;
    padding: 10px;
    outline: none;
    border: none;
    border-bottom: 3px solid darkcyan;
  }
  .input-field:focus {
    border: 2px solid darkcyan;
  }

  button {
    background-color:darkseagreen;
    color: white;
    padding: 15px 20px;
    border: none;
    cursor: pointer;
    width: 100%;
    opacity: 0.9;
  }
  .bttn:hover {
    opacity: 1;
    background-color: darkcyan;
  }  
  a:hover{
    color: blueviolet;
}
.icon:hover{
    background-color: darkcyan;
}
</style>
