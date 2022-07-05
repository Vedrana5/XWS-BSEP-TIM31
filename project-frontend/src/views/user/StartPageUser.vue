<template>
  <div>
        <button @click="Update()">Update profile info</button>
        <button @click="ResetPassword()">Reset Password</button>
         <button @click="CreatePost()">Create post</button>
          <div><button @click="Find()">Find public user</button></div>
  </div>
</template>


<script>
import axios from 'axios'

export default {
  name: "StartPageUser",
   data: () => ({
    err: "" ,
    Username: "",
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
        TypeOfProfile:"PUBLIC",
        TypeOfUser:"REGISTERED_USER",
        Biography:"",
        WorkExperience:"",
        Education:"",
        Skills:"",
        Interest:"",
      },
  }),
  methods: {
    async getUser() {
               const headers ={
            Authorization: "Bearer " + this.token,
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json',
          }
  axios.get("http://localhost:8089/findByUsername/"+ this.Username, {
      headers}
    )
    .then((response) => {

      //console.log(this.response.FirstName)
      console.log(response.data.gender)
      this.newUser.Username = response.data.username
      this.newUser.Email = response.data.email
      this.newUser.PhoneNumber = response.data.phoneNumber
      this.newUser.FirstName = response.data.firstName
      this.newUser.LastName = response.data.lastName
      if(response.data.gender == 1) {
      this.newUser.Gender = "MALE"
      }else if (response.data.gender == 2) {
        this.newUser.Gender = "FEMALE"
      }else {
        this.newUser.Gender = "OTHER"
      }
      this.newUser.TypeOfProfile = response.data.typeOfProfile

      this.newUser.TypeOfUser = response.data.typeOfUser
      this.newUser.Biography = response.data.biography
      this.newUser.WorkExperience = response.data.workExperience
      this.newUser.Education = response.data.education
      this.newUser.Skills = response.data.skills
      this.newUser.Interest = response.data.interest
      this.newUser.Question = response.data.question
        })

    },
      async Update() {
          this.$router.push({ name: "UpdateInfoView" });
      },
      async CreatePost() {
          this.$router.push({ name: "CreatePost" });
      },
  async Find() {
    this.$router.push({ name: "FindPublicUserByLogUser" });
  },

  },
  async created() {
  this.Username = localStorage.getItem("username"); 
  this.token = localStorage.getItem("token")
  console.log("username je"+ this.Username);
  console.log("token je"+ this.token);
  this.type = localStorage.getItem("userType")
  console.log("type je"+ this.type);
  // await this.getUser(this.Username);
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
