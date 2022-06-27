<template>
  <div>
      <input v-model="validationToken.Code"/>
      <button @click.prevent="Validations()">Valid</button>
  </div>
  <div v-if=" this.counter === 1">

              <form class="myForm" name="myForm">
      <div>
        <label for="password">New password</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="Password" required>
      </div>
            <div>
        <label for="password"> New password again</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="passwordAgain" required>
      </div>
            <div><button type="submit" @click="setPassword()">Set new password</button></div>
         </form>
  </div>
</template>

<script>
import Swal from 'sweetalert2';
import axios from 'axios';
export default {
  name: "ConfirmResetPassword",
  data: () => ({
    counter : 0,
    confirmationToken:"",
    id: "",
    token:"",
    d: Date(),
        passwordAgain:"",
        Password:"",
                  newUser: {
                    ID:0,
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
      validationToken: {
          UserId: 0,
          IsValid:null,
          IsUsed: null,
          Code: null,   
          ExpiredDate: null,
      }
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
      this.confirmationToken = hrefPaths[5];
      alert("Id je"+ this.id + "Kod je" + this.confirmationToken)
      this.getUserr();
      this.getCode();

    },
    async getUserr() {
  axios.get("http://localhost:8089/findById/"+ this.id
    )
    .then((response) => {
      console.log(response.data.gender)
      this.newUser.ID = response.data.id
      this.newUser.Username = response.data.username
      this.newUser.Email = response.data.email
      this.newUser.FirstName = response.data.firstName
      this.newUser.LastName = response.data.lastName
      console.log(this.newUser.Username)
        })
    },

        async getCode() {
               const headers ={
            Authorization: "Bearer " + this.token,
            'Content-Type': 'application/json;charset=UTF-8',
            Accept: 'application/json',
          }
  axios.get("http://localhost:8089/findTokenByCode/"+ this.confirmationToken, {
      headers}
    )
    .then((response) => {
      console.log(response.data) 
      this.validationToken.Code = response.data.code;
      this.validationToken.UserId = response.data.user_id
      this.validationToken.IsUsed = response.data.is_used
      this.validationToken.IsValid = response.data.is_valid
      this.validationToken.ExpiredDate = response.data.expired_time
        })

    

    },
    async Validations() {
              console.log(this.newUser.ID);
        console.log(this.validationToken.UserId)
        console.log(this.validationToken.IsValid);
        console.log(this.validationToken.ExpiredDate);
        console.log(this.d)
      if (this.newUser.ID != this.validationToken.UserId) {
            new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Ovaj kod nije tvoj!',
           });
      }
      else if(this.validationToken.IsUsed) {
            new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Ovaj kod je vec koristen!',
           });
      

      // }else if(this.validationToken.ExpiredDate < this.d) {
      //       new Swal({
      //        title:"Nije uspesno",
      //        type: "warning",
      //        text:'Ovaj kod je istekao!',
      //      });

      }else {

        this.counter = 1;
      }


    },
    async setPassword() {
        if (!this.validPassword()) {
            alert("Password isn't valid!")
        }
        else {
 
  axios.post("http://localhost:8089/resettPassword",{
                ID: this.newUser.ID,
                Password: this.Password,
                Code: this.validationToken.Code
            }
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
input {
  width: 100%;
}
</style>