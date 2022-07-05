<template>
  <div>
<button @click="GoBack()">Go back</button>
      <input
        class="form-control mr-sm-2"
        type="text"
        placeholder="Username"
        v-model="username"
      />
      <button class="btn btn-success" type="submit" @click="Search(username)">
        Search
      </button>

        <div class="containerInfo">
    <div class="tab-pane container active">
      <div class="row-boats" v-for="(user, index) in users" :key="index">
        <div class="col-info">
          <h4 style="width: 600px" class="text">
            Username: {{user.Username}}
          </h4>
          <h4 style="width: 600px" class="text">Firstname: {{user.FirstName}}</h4>
          <h4 style="width: 600px" class="text">
            Lastname: {{user.LastName}}
          </h4>
          <img src=""/>
          <button  class="btn btn-success" @click="FindPosts(user.Username)">Go to profile</button>
        </div>
      </div>
    </div>
  </div>

<div v-if="this.counter==1" >
                      <div v-for="(Post, index) in posts" :key="index">
                        <h4 style="width: 600px" class="text">
                          Date: {{Post.DatePosted}}
                        </h4>
                       <h4 style="width: 600px" class="text">
                          Text: {{Post.PostText}}
                        </h4>
                        <h4 style="width: 600px" class="text">Photo:</h4>
                        <img v-bind:src="'data:image/jpeg;base64,'+Post.ImagePaths"/>
                        <h4 style="width: 600px" class="text">
                          Likes({{Post.LikesNumber}})
                        </h4>
                        
                        <h4 style="width: 600px" class="text">
                          Dislikes({{Post.DislikesNumber}})
                        </h4>
                        <h4 style="width: 600px" class="text">
                          Comments({{Post.CommentsNumber}})
                        </h4>
                      </div>

</div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "FindPublicUser",
  data() {
    return {
      username:"",
      users:"",
      posts:"",
       Post:{PostText:"", ImagePaths:null},
       counter:0,
      user: {
        username:"",
       firstName:"",
        lastName:"",
      },
    };
  },


  methods: {
    async GoBack() {
      this.$router.push({ name: "StartPageView" });
    },
    async Search(username) {
      this.username= username;

     await this.FindUser(username);
    },
  async FindUser(username) {
    this.username= username;
          axios.get("http://localhost:9090/publicUser/"+this.username)
      .then (response => { 
          this.users = response.data.users;
      })  
  },
    async FindPosts(username) {
      this.counter=1;
        this.username= username;
        console.log("username je"+this.username);
          axios.get("http://localhost:9090/post/user/"+this.username)
      .then (response => { 
          this.posts = response.data.Posts;
          console.log("Odgovor je"+response.data.Posts)
      })
  },
  async findOneUser(username){
        this.username= username;
          axios.get("http://localhost:9090/publicUser/"+this.username)
      .then (response => { 
          this.user = response.data.users[0];
      })  
      await this.FindPosts(username);
       console.log("Broj postova je"+this.posts.length)
  }
  }
};
</script>

<style scoped>
.row-boats {
  display: flex;
}
.text {
  text-align: left;
  font-size: 15px;
  font-weight: 600;
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-info {
  margin-left: 30px;
  margin-top: 10px;
}
.row {
  width: 660%;
  padding-left: 60px;
  height: 10%;
  margin-top: 100px;
  margin-left: 5000px;
  margin-right: 500px;
  opacity: 1.2;
  border-radius: 10px;
  align-content: center;
  background-color: rgba(255, 255, 255, 0.288);
  box-shadow: 0 5px 4px rgb(0 0 0 / 30%), 0 1px 1px rgb(0 0 0 / 22%);
}
/* kartica u okviru stavke menija  */
.containerInfo {
  width: 97%;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 10px;
  opacity: 1;
  border-radius: 10px;
  align-content: center;
  position: relative;
  background-color: fff;
  box-shadow: 0 19px 40px rgb(0 0 0 / 30%), 0 15px 12px rgb(0 0 0 / 22%);
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
</style>