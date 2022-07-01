<template>
  <div>
     <h1 id="heading1">Update info Here</h1>
       <img src="https://www.kindpng.com/picc/m/273-2738790_login-login-logo-hd-png-download.png" style="width:200px; height:150px;">
     <div>
 <form class="myForm" name="myForm">
      <div>
        <label for="username">Username</label>
        <input name="username" class="input-field" placeholder="username" v-model="newUser.Username" required>
      </div>
      <div>
        <label for="firstName">First name</label>
        <input name="firstName" class="input-field" placeholder="first name" v-model="newUser.FirstName" required>
      </div>
      <div>
        <label for="lastName">Last name</label>
        <input name="lastName" class="input-field" placeholder="last name" v-model="newUser.LastName" required>
      </div>
      <div>
        <label for="email">Email</label>
        <input name="email" class="input-field"  placeholder="email" v-model="newUser.Email" required>
      </div>
      <div>
        <label for="phoneNumber">Phone number</label>
        <input name="phoneNumber" class="input-field" placeholder="phoneNumber"  v-model="newUser.PhoneNumber" required>
      </div>
       <div>
        <label for="date">Date of birth</label>
        <input type="date" class="input-field" name="date" placeholder="date"  
        v-model="newUser.DateOfBirth" required>
      </div>
             <div>
        <label for="biography">Biography</label>
        <input name="biography" class="input-field" placeholder="biography"  v-model="newUser.Biography" required>
      </div>
             <div>
        <label for="workExperience">Work experience</label>
        <input name="workExperience" class="input-field" placeholder="workExperience"  v-model="newUser.WorkExperience" required>
      </div>
      <div>
        <label for="gender">Gender</label>
        <select class="input-field" v-model="newUser.Gender" required>
          <option value="MALE">MALE</option>
          <option value="FEMALE">FEMALE</option>
          <option value="OTHER">OTHER</option>
        </select>
      </div>
                   <div>
        <label for="education">Education</label>
        <input class="input-field"  name="education" placeholder="education"  v-model="newUser.Education" required>
      </div>
                   <div>
        <label for="skills">Skills</label>
        <input class="input-field" name="skills" placeholder="skills"  v-model="newUser.Skills" required>
      </div>
                   <div>
        <label for="interest">Interest</label>
        <input class="input-field" name="interest" placeholder="interest"  v-model="newUser.Interest" required>
      </div>
      <div><button type="submit" @click="Update()">Update</button></div>
    </form>
     </div>
  </div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2';
import moment from 'moment';
export default {
  name: "UpdateInfoView",
 data() {
    return {
      id:"",
      type:"",
      oldUsername:"",
        Username: "",
        passwordAgain:"",
        Password:"",
        Email:"",
        PhoneNumber:"",
        FirstName:"",
        LastName:"",
        DateOfBirth:"",
        Gender:"",
        Biography:"",
        WorkExperience:"",
        Education:"",
        Skills:"",
        Interest:"",
        Question:"",
        Answer:"",
        TypeOfProfile:"",
        TypeOfUser:"",
        errorMessage:"",
        token:"",
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
    }
    },
    methods: {
      async Update() {
 if (
  
        !this.validBiography() ||
        !this.validPhoneNumber() ||
        !this.validLastName() ||
        !this.validUsername() ||
        !this.validFirstName()
      ) {
      alert("Fill in the fields in the right way!");
      } else {
        this.id = localStorage.getItem("userId");
        this.oldUsername = localStorage.getItem("username");
        

axios.post("http://localhost:9090/Euser",{          
      ID: this.id,
       OldUsername: this.oldUsername,     
       Username : this.newUser.Username, 
       Password : this.newUser.Password,
       Email : this.newUser.Email,
       PhoneNumber : this.newUser.PhoneNumber,
       DateOfBirth: this.newUser.DateOfBirth,
       FirstName : this.newUser.FirstName,
       LastName : this.newUser.LastName,
       Gender : this.newUser.Gender,
       TypeOfUser: this.newUser.TypeOfUser,
       TypeOfProfile : this.newUser.TypeOfProfile,
       Biography : this.newUser.Biography,
       WorkExperience : this.newUser.WorkExperience,
       Education : this.newUser.Education,
       Skills : this.newUser.Skills,
       Interest : this.newUser.Interest,}
    ).then((res) => {
          console.log(res);
           new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Izmena profila je uspela!',
           });
           localStorage.setItem("username", this.newUser.Username);
        })
        .catch((error) => {
          console.log(error.response.status);
          if(error.response.status == 417) {
           new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Postoji vec korisnik sa tim email-om!',
           });
          }else if (error.response.status == 409) {
           new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Postoji vec korisnik sa tim username-om!',
           });
          }
        });
      }
      },

      validQuestion() {
      if (this.newUser.Question.length < 1) {
        alert("Your question should contain at least 1 character!");
        return false;
      } else if (this.newUser.Question.length > 70) {
        alert("Your question shouldn't contain more than 70 characters!");
        return false;
      }
      return true;
    },
    validAnswer() {
      if (this.newUser.Answer.length < 1) {
        alert("Your answer should contain at least 1 character!");
        return false;
      } else if (this.newUser.Answer.length > 70) {
        alert("Your answer shouldn't contain more than 70 characters!");
        return false;
      }
      return true;
    },
      validBiography() {
      if (this.newUser.Biography.length < 1) {
        alert("Your biography should contain at least 1 character!");
        return false;
      } else if (this.newUser.Biography.length > 100) {
        alert("Your biography shouldn't contain more than 100 characters!");
        return false;
      }
      return true;
    },
     validPhoneNumber() {
      if (this.newUser.PhoneNumber.match(/[a-zA-Z]/g)) {
        alert("Your phone number shouldn't contain letters.");
        return false;
      } else if (this.newUser.PhoneNumber.match(/[ ]/g)) {
        alert("Your phone number shouldn't contain spaces!");
        return false;
      } else if (
        !/^[+]*[(]{0,1}[0-9]{1,3}[)]{0,1}[-\\s./0-9]*$/.test(this.newUser.PhoneNumber)
      ) {
        alert("Your phone number is not in right form!");
        return false;
      }
      return true;
    },
    validLastName() {
      if (this.newUser.LastName.length < 2) {
        alert("Your last name should contain at least 2 characters!");
        return false;
      } else if (this.newUser.LastName.length > 35) {
        alert("Your last name shouldn't contain more than 35 characters!");
        return false;
      } else if (this.newUser.LastName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
        alert("Your last name shouldn't contain special characters.");
        return false;
      } else if (this.newUser.LastName.match(/[ ]/g)) {
        alert("Your last name shouldn't contain spaces!");
        return false;
      } else if (this.newUser.LastName.match(/\d/g)) {
        alert("Your last name shouldn't contain numbers!");
        return false;
      } else if (!/^[A-Z][a-z]+$/.test(this.newUser.LastName)) {
        alert("Your last name needs to have one upper letter at the start!");
        return false;
      }
      return true;
    },
    validUsername() {
      if (this.newUser.Username.length < 1) {
        alert("Your username should contain at least 1 character!");
        return false;
      } else if (this.newUser.Username.length > 20) {
        alert("Your username shouldn't contain more than 20 characters!");
        return false;
      } else if (this.newUser.Username.match(/[!@#$%^&*'<>+/\\"]/g)) {
        alert("Your username shouldn't contain special characters.");
        return false;
      }
      return true;
    },
     validEmail() {
      if (
        !/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(
          this.newUser.Email
        )
      ) {
        alert("You have entered an invalid email address!");
        return false;
      } else if (this.newUser.Email.length > 35) {
        alert("Your email shouldn't contain more than 35 characters!");
        return false;
      }
      return true;
    },

    validFirstName() {
      if (this.newUser.FirstName.length < 2) {
        alert("Your first name should contain at least 2 characters!");
        return false;
      } else if (this.newUser.FirstName.length > 20) {
        alert("Your first name shouldn't contain more than 20 characters!");
        return false;
      } else if (this.newUser.FirstName.match(/[!@#$%^&*.,:'<>+-/\\"]/g)) {
        alert("Your first name shouldn't contain special characters.");
        return false;
      } else if (this.newUser.FirstName.match(/[ ]/g)) {
        alert("Your first name shouldn't contain spaces!");
        return false;
      } else if (this.newUser.FirstName.match(/\d/g)) {
        alert("Your first name shouldn't contain numbers!");
        return false;
      } else if (!/^[A-Z][a-z]+$/.test(this.newUser.FirstName)) {
        alert("Your first name needs to have one upper letter at the start!");
        return false;
      }
      return true;
    },
    },
async created() {    
  this.Username = localStorage.getItem("username"); 
  this.token = localStorage.getItem("token")
  this.type = localStorage.getItem("userType")
  console.log("type je"+ this.type);

       console.log()
  axios.get("http://localhost:9090/user/"+this.Username
    )
    .then(response => {
      console.log("datum je "+ response.data.user.DateOfBirth.substring(0,10) );
     this.newUser.Username = response.data.user.Username
      this.newUser.Email = response.data.user.Email
      this.newUser.PhoneNumber = response.data.user.PhoneNumber
      this.newUser.FirstName = response.data.user.FirstName
      this.newUser.LastName = response.data.user.LastName
      this.newUser.Gender = response.data.user.Gender
      this.newUser.TypeOfProfile = response.data.user.TypeOfProfile
      this.newUser.DateOfBirth = moment(response.data.user.DateOfBirth.substring(0,10)).format('YYYY-MM-DD')
      this.newUser.TypeOfUser = response.data.user.TypeOfUser
      this.newUser.Biography = response.data.user.Biography
      this.newUser.WorkExperience = response.data.user.WorkExperience
      this.newUser.Education = response.data.user.Education
      this.newUser.Skills = response.data.user.Skills
      this.newUser.Interest = response.data.user.Interest
      this.newUser.Question = response.data.user.Question
        })
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
