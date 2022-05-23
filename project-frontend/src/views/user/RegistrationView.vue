<template>
  <div>
     <h1 id="heading1">Registration Here</h1>
       <img src="https://www.kindpng.com/picc/m/273-2738790_login-login-logo-hd-png-download.png" style="width:200px; height:150px;">
     <div>
 <form @submit.prevent="login">
      <div>
        <label for="username">Username</label>
        <input name="username" placeholder="username">
      </div>
      <div>
        <label for="password">Password</label>
        <input name="password" placeholder="password" type="password">
      </div>
      <div>
        <label for="firstName">First name</label>
        <input name="firstName" placeholder="first name">
      </div>
      <div>
        <label for="lastName">Last name</label>
        <input name="lastName"  placeholder="last name">
      </div>
      <div>
        <label for="email">Email</label>
        <input name="email" placeholder="email">
      </div>
      <div>
        <label for="phoneNumber">Phone number</label>
        <input name="phoneNumber" placeholder="phoneNumber">
      </div>
       <div>
        <label for="date">Date of birth</label>
        <input type="date" name="date" placeholder="date">
      </div>
             <div>
        <label for="biography">Biography</label>
        <input name="biography" placeholder="biography">
      </div>
             <div>
        <label for="workExperience">Work experience</label>
        <input name="workExperience" placeholder="workExperience">
      </div>

                   <div>
        <label for="education">Education</label>
        <input name="education" placeholder="education">
      </div>
                   <div>
        <label for="skills">Skills</label>
        <input name="skills" placeholder="skills">
      </div>
                   <div>
        <label for="interest">Interest</label>
        <input name="interest" placeholder="interest">
      </div>
                   <div>
        <label for="question">Question</label>
        <input name="question" placeholder="question">
      </div>
                   <div>
        <label for="answer">Answer</label>
        <input name="answer" placeholder="answer">
      </div>
      <input type="submit" value="Register">
    </form>
     </div>
  </div>
</template>

<script>
export default {
  name: "RegistrationView",
  data() {
    return {

    };
  },

  methods: {
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
  .a {
          width: 100%;
    padding: 10px;
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
