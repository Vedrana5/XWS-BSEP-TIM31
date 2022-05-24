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
        <label for="password">Password</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="newUser.Password" required>
      </div>
            <div>
        <label for="password">Password again</label>
        <input name="password" class="input-field" placeholder="password" type="password" v-model="this.passwordAgain" required>
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
        <input type="date" class="input-field" name="date" placeholder="date"  v-model="newUser.DateOfBirth" required>
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
          <option>MALE</option>
          <option>FEMALE</option>
          <option>OTHER</option>
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
                   <div>
        <label for="question">Question</label>
        <input class="input-field" name="question" placeholder="question"  v-model="newUser.Question" required>
      </div>
                   <div>
        <label for="answer">Answer</label>
        <input class="input-field" name="answer" placeholder="answer"  v-model="newUser.Answer" required>
      </div>
      <div><button type="submit" @click="Register()">Registration</button></div>
    </form>
     </div>
  </div>
</template>

<script>
export default {
  name: "UpdateInfoView",
 data() {
    return {
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
        Question:"",
        Answer:"",
      },
    }
    }
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
