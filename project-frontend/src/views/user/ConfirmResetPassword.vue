<template>
  <div >
          <form class="myForm" name="myForm">
      <div>
        <label for="password">Password</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="Password" required>
      </div>
            <div>
        <label for="password">Password again</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="passwordAgain" required>
      </div>
            <div><button type="submit" @click="setPassword()">Set new password</button></div>
         </form>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "ConfirmResetPassword",
  data: () => ({
    confirmationToken:"",
    id: "",
    token:"",
        passwordAgain:"",
        Password:"",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      var hrefPath = window.location.href;
      var hrefPaths = [];
      hrefPaths=  hrefPath.split('/');
      this.id = hrefPaths[4];
      alert(this.id)
    },
    async setPassword() {
        if (!this.validPassword()) {
            alert("Password isn't valid!")
        }
        else {
            this.token = localStorage.getItem("token")
         const headers ={
            Authorization: "Bearer " + this.token,
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json',
          }  
  axios.post("http://localhost:8089/resetPassword",{
                ID: this.id,
                Password: this.Password,
                Token: this.token
            }, {
      headers}
    )
        .then((res) => {
          console.log(res);
          this.message = "You have successfully change your password! You can log in on system!"
        })
        .catch((err) => {
          console.log(err);
          this.message = "There is a problem with changing your password! Please, contact system admin!"
        });
        }
    },
        validPassword() {
      if (this.Password.length < 10) {
        alert("Your password should contain at least 10 character!");
        return false;
      } else if (this.Password.length > 30) {
        alert("Your password shouldn't contain more than 30 characters!");
        return false;
      } else if (!this.Password.match(/[a-z]/g)) {
        alert("Your password should contain at least one small letter.");
        this.passwordAgain = "";
        return false;
      } else if (!this.Password.match(/[A-Z]/g)) {
        alert("Your password should contain at least one big letter.");
        this.passwordAgain = "";
        return false;
      } else if (!this.Password.match(/[0-9]/g)) {
        alert("Your password should contain at least one number.");
        this.passwordAgain = "";
        return false;
      } else if (!this.Password.match(/[!@#$%^&*.,:'+-/\\"]/g)) {
        alert(
          "Your password should contain at least one special character (other than <>)."
        );
        return false;
      } else if (this.Password.match(/[<>]/g)) {
        alert("Your password shouldn't contain special character < or >.");
        return false;
      } else if (this.Password.match(/[ ]/g)) {
        alert("Your password shouldn't contain spaces!");
        return false;
      } else if (this.Password !== this.passwordAgain) {
        alert("Passwords don't match !!!");
        this.passwordAgain = "";
        return false;
      }
      return true;
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