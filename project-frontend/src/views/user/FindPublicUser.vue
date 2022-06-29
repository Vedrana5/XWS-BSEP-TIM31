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
            Username: {{user.username}}
          </h4>
          <h4 style="width: 600px" class="text">Firstname: {{user.firstName}}</h4>
          <h4 style="width: 600px" class="text">
            Lastname: {{user.lastName}}
          </h4>
          <button  class="btn btn-success">Go to profile</button>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: "FindPublicUser",
  data() {
    return {
      username:"",
      users:"",
      user: {
        username:"",
       firstName:"",
        lastName:""
      },
    };
  },


  methods: {
    async GoBack() {
      this.$router.push({ name: "StartPageView" });
    },
    async Search(username) {
      this.username= username;

      this.users = await this.FindUser(username);
    },
  async FindUser(username) {
    this.username= username;
      const res = await fetch("http://localhost:8089/findPublic/"+this.username);
      const data = await res.json();
      console.log(data);
      return data;
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