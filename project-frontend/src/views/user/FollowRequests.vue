<template>
  <div>
    <div>   <button @click="GoBack()">Go back</button></div>
<table class="styled-table">
    <thead>
        <tr>
            <th>Username of user who wants to follow you</th> 
            <th></th> 
            <th></th>        
        </tr>
    </thead>
    <tbody>
          <tr v-for="(request, index) in requests" :key="index">
                  <td>{{request.FirstUsername}}</td>
                  <td><button class="btn btn-success btn-block">Confirm</button></td>
                  <td><button class="btn btn-success btn-block">Reject</button></td>
           </tr> 
    </tbody>
</table>

  </div>
</template>


<script>
import axios from 'axios'

export default {
  name: "FollowRequests",
   data: () => ({
    requests:"",
    request:{Id:0,FirstUsername:"",SecondUsername:"",IsConfirmed:false,IsDeleted:false},
    username:""
  }),
  methods: {
    async GoBack() {
        this.$router.push({ name: "StartPageUser" });
    },
    async GetRequests() {
    this.username = localStorage.getItem("username");
    axios.get("http://localhost:9090/findConnection/"+this.username)
    .then (response => { 
          this.requests = response.data.connection;
    }) 

    }
  },
  async created() {
    await this.GetRequests();
  }
  
}
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
.styled-table {
    border-collapse: collapse;
    margin: 25px 20px;
    font-size: 0.9em;
    font-family: sans-serif;
    min-width: 400px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
}
h1 {
    margin: 25px 20px;
    font-family: sans-serif;
}
.styled-table thead tr {
    background-color: #4caf50;
    color: #ffffff;
    text-align: left;
}
button {
  background-color: #4caf50; /* Green */
}
.styled-table th,
.styled-table td {
    padding: 12px 15px;
}
.icon:hover{
    background-color: darkcyan;
}
</style>
