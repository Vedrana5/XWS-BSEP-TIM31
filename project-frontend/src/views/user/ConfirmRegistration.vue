<template>
  <div >
    <h1 class="center">{{ message }}</h1>
    <button @click="Login()">Login</button>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "ConfirmRegistration",
  data: () => ({
    confirmationToken:"",
    userId: "",
    message: ""
  }),
  mounted() {
    this.init();
  },
  methods: {
    Login() {
 this.$router.push({ name: "LoginView" });
    },
    init() {
      console.log(" HIIIIIIIIIIII");
      var hrefPath = window.location.href;
      var hrefPaths = [];
      hrefPaths=  hrefPath.split('/');
      this.confirmationToken = hrefPaths[4];
      this.userId = hrefPaths[5];
      alert(this.confirmationToken+"/"+this.userId)
    const headers ={
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json',
          }
  console.log(this.Username);
  axios.post("http://localhost:8089/confirmRegistration",{          
          confirmation_token: this.confirmationToken,
          user_id: this.userId,}, {
      headers}
    )
        .then((res) => {
          console.log(res);
          this.message = "You have successfully verified your account! You can log in on system!"
        })
        .catch((err) => {
          console.log(err);
          this.message = "Your token is invalid or expiried! Please, contact system admin!"
        });
    },
    redirectToLogIn() {
      window.location.href = "https://localhost:8082/LoginView";
    },
  },
};
</script>

<style scoped>
.helloMessage {
  font-weight: bolder;
  font-size: 20px;
  height: 50px;
}

.center {
  margin-top: 10%;
  padding: 10px;
  text-align: center;
}

#certificateCard {
  margin-top: 5%;
  width: 70%;
  height: 760px;
  overflow-y: scroll;
}
</style>