<template>
  <div>
     <h1 id="heading1">Login Here</h1>
       <img src="https://www.kindpng.com/picc/m/273-2738790_login-login-logo-hd-png-download.png" style="width:200px; height:150px;">
     <div>
         <form class="myForm" name="myForm">
             <div class="class1">
                 <input  placeholder="Username" name="username" class="input-field" v-model="Username" required>
             </div>
             <div class="class1">
           <input v-if="showPassword" type="text" class="input-field" v-model="Password" />
           <input v-else type="password" class="input-field" v-model="Password">
           <button class="class2" @click.prevent="Show()">Show</button></div>
            <div><button type="submit" @click.prevent="Login()">Login</button>
            <button class="class2" @click.prevent="ForgotPassword()">Forgot password?</button>
             <button class="class2" @click.prevent="PasswordlessLogin()">Passwordless login</button></div>
         </form>
     </div>
  </div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2';

export default {
  name: "LoginView",
  data: () => ({
    err: "" ,
    showPassword: false,
    Username: "",
    Password: "",
    Question: "",
    Answer: "",
  }),

  methods: {
async Login() {          
      axios.post("http://localhost:9090/user",{           
          Username: this.Username,
          Password: this.Password,
       })
      .then (response => { 
        console.log("ID JE"+response.data.loginUserResponse)
          localStorage.setItem("username", this.Username);
          localStorage.setItem("token", response.data.loginUserResponse.Token);
          localStorage.setItem("userId", response.data.loginUserResponse.ID);
          localStorage.setItem("userType", response.data.loginUserResponse.TypeOfUser);
                      new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Uspesno ste ulogovani!',
           });     
            this.$router.push({ name: "StartPageUser" });
      })      .catch(function (error) {
            console.log(error.response.status)
            new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Neuspesno logovanje!',
           }); 
         
        });
         // this.$router.go(0);
     
  },
  async ForgotPassword() {
       this.$router.push({ name: "ResetPassword" });
  },
  async PasswordlessLogin() {
      this.$router.push({ name: "PasswordlessLogin" });
  },
  async Show() {
    if (this.showPassword== true) {
      this.showPassword = false;
    } else {
      this.showPassword = true;
    }
  },
    
},
};
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
    .class2 {
    
    background-color:darkseagreen;
    color: white;
    padding: 15px 20px;
    border: none;
    cursor: pointer;
    width: 20%;
    
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
