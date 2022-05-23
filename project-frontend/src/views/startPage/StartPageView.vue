<template>
  <div>
     <h1 id="heading1">Welcome to our site</h1>
       <img src="https://rietveld-ict.nl/wp-content/uploads/2014/01/users.png" style="width:200px; height:150px;">
     <div>
         <div> <button @click="Login()">Login</button></div>
            <div><button @click="Registration()">Registration</button></div>
            <div><button @click="Find()">Find public user</button></div>
     </div>
  </div>
</template>

<script>
export default {
  name: "StartPageView",
    methods: {
async Login() {
    this.$router.push({ name: "LoginView" });
  },
async Registration() {
    this.$router.push({ name: "RegistrationView" });
  },
  async Find() {
    this.$router.push({ name: "FindPublicUser" });
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
#router-link {
  display: block;
  margin-left: 60px;
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
    width: 15%;
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
