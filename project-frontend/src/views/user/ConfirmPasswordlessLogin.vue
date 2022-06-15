<template>
  <div>
      <input v-model="validationToken.Code"/>
      <button @click.prevent="Validations()">Valid</button>
  </div>
  <div v-if=" this.counter === 1">
            <div><button type="submit" @click="Login()">Login</button></div>
  </div>
</template>

<script>
import Swal from 'sweetalert2';
import axios from 'axios';
export default {
  name: "ConfirmPasswordlessLogin",
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
      console.log(response.data)
      this.newUser.ID = response.data.id
      this.newUser.Username = response.data.username
      this.newUser.Email = response.data.email
      this.newUser.FirstName = response.data.firstName
      this.newUser.LastName = response.data.lastName
      this.newUser.Password = response.data.password
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
    async Login() {
      axios.post("http://localhost:8089/loginPasswordless",{           
          Username: this.newUser.Username,
          Code:  this.confirmationToken
       })
      .then (response => { 
        console.log(response.data.Token)
          localStorage.setItem("username", this.newUser.Username);
          localStorage.setItem("token", response.data.Token);
          localStorage.setItem("userId", response.data.ID);
          localStorage.setItem("userType", response.data.TypeOfUser);
        
      })
      .catch((err) => {
          console.log(err);
        });
           this.$router.push({ name: "StartPageUser" });
    }

  },
};
</script>

<style scoped>
input {
  width: 100%;
}
</style>