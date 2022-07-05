<template>
  <div>
     <h1 id="heading1">Create post</h1>
     <div>
         <form class="myForm" name="myForm">
             <div class="class1">
                 <input  placeholder="Type text" name="utextsername" class="input-field" v-model="Post.Text" required>
             </div>
              <div class="upload-images">
                    <input type="file"  @change="imageAdded"/>
                </div> <br/>
                <div v-if="imagesFrontend" class="images-preview">
                        <div v-for="image in imagesFrontend" :key="image">
                            <img :src="image" />
                        </div>
                </div>
            <div><button type="submit" @click.prevent="CreatePost()">Create post</button></div>
         </form>
     </div>
  </div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2';

export default {
  name: "LoginView",
  data: () => ({
    err: "" ,
    showPassword: false,
    username: "",
    Password: "",
    Question: "",
    Answer: "",
    Post:{Text:"", Image:null},
    imagesFrontend:[]
  }),

  methods: {
    imageAdded(e) {
           
                const file = e.target.files[0];  
                console.log(file)        
                this.createBase64Image(file);
                this.imagesFrontend.push(URL.createObjectURL(file));
                
            },
            createBase64Image(file){
                const reader= new FileReader();
                console.log("OVDE SAM1");
                reader.onload = (e) =>{
                    let img = e.target.result;
                     console.log(img)  
                    this.Post.Image=img;
                    console.log("OVDE SAM2");
                }
                reader.readAsDataURL(file);
            },
async CreatePost() {  
    this.username = localStorage.getItem("username");        
      axios.post("http://localhost:9090/post",{ 
          Username: this.username,          
          PostText: this.Post.Text,
          ImagePaths: this.Post.Image,
       })
      .then (response => { 
        console.log("RESPONSE JE"+response.data)
                      new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Uspesno ste dodali post!',
           });     
            this.$router.push({ name: "StartPageUser" });
      })
    
     
  },
  async GoBack() {
       this.$router.push({ name: "StartPageUser" });
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
