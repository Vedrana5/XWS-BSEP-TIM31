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
                <input type="showPassword ? 'text' : 'password'" placeholder="Password" name="password" class="input-field" v-model="Password"           :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
           required></div>
                        <div class="class1">
                 <input  placeholder="Question" name="question" class="input-field" v-model="Question" required>
             </div>
                          <div class="class1">
                 <input  placeholder="Answer" name="answer" class="input-field" v-model="Answer" required>
             </div>
            <div><button type="submit" @click="Login()">Login</button></div>
         </form>
     </div>
  </div>
</template>

<script>
export default {
  name: "LoginView",
  data: () => ({
    showPassword: true,
    Username: "",
    Password: "",
    Question: "",
    Answer: "",
  }),
  computed: {
    user() {
      return { Username: this.Username, Password: this.Password };
    },
  },

  methods: {
async Login() {
        fetch("http://localhost:8089/login",{
          method:"POST",
          body: JSON.stringify({
          Username: this.Username,
          Password: this.Password,
          Question: this.Question,
          Answer: this.Answer,
          }),
        headers: {
          'Content-type':'application/json; charset=UTF-8',
        },
        })
        .then((response) => response.json())
        .then((json) => console.log(json))
        .catch(error => {
      this.errorMessage = error;
       alert("Login failed!");
    });
      
  }
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
