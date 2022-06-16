<template>
      <div>
     <h1 id="heading1">RESET PASSWORD HERE</h1>
       <img src="https://www.kindpng.com/picc/m/273-2738790_login-login-logo-hd-png-download.png" style="width:200px; height:150px;">
     <div>
 <form class="myForm" name="myForm">
      <div>
        <label for="email">Email</label>
        <input name="email" class="input-field"  placeholder="email" v-model="Email" required>
      </div>
      <div> <button type="submit" @click="SendLink(Email)">Send link</button></div>
    </form>
     </div>
  </div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2';

export default {
  name: "PasswordlessLogin",
  data: () => ({
    err: "" ,
    Email: "",
  }),

  methods: {
async SendLink(email) {          
       this.Email= email;
      axios.post("http://localhost:8089/linkForPasswordless/"+this.Email,{           
       })
         .catch(function (error) {
            console.log(error.response.status)
            new Swal({
             title:"Nije uspesno",
             type: "warning",
             text:'Niste uneli email!',
           }); 
         
        });
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
