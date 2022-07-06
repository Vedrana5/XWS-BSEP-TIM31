<template>
    <div>
        <button @click="GoBack()">Go back</button>
        <h1 id="heading1">Create Job Offer</h1>
        <div>
            <form class="myForm" name="myForm">
                <div class="class1">
                    <label>Position</label>
                    <input  placeholder="Type text" name="position" class="input-field" v-model="Position" required>
                </div>
                <div class="class1">
                    <label>Job Description</label>
                    <input  placeholder="Type text" name="position" class="input-field" v-model="JobDescription" required>
                </div>
                    <button data-target="#addLink" data-toggle="modal">Add Requirements</button>         
                    <div v-for="Requirement in Requirements" :key="Requirement">
                    <label>{{Requirement}}</label>
                    </div>
                    
                <div><button type="submit" @click.prevent="CreateJobOffer()">Create job offer</button></div>
            </form>
        </div>
    </div>

    <div class="modal fade" id="addLink" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-body" style="padding: 15px 50px">
                    <form role="form" @submit.prevent="AddRequirement">
                        <div class="form-group">
                            <label><span class="glyphicon glyphicon-eye-open"></span>Requirement:</label>
                            <input type="text" v-model="Requirement"/>
                        </div>
                        <button type="submit" class="btn btn-success btn-block">
                            <span></span> Add Requirement
                        </button>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" data-dismiss="modal">Odustani</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import Swal from 'sweetalert2';
export default {
  name: "CreateJobOffer",
  data() {
    return {
        Position:"",
        JobDescription:"",
        Requirements:[],
        Requirement:"",
        username:""
    };
  },
methods:{
    async GoBack() {
        this.$router.push({ name: "StartPageUser" });
    },
    async AddRequirement() {
        console.log(this.Requirement)
        this.Requirements.push(this.Requirement);
    },
    async CreateJobOffer() {
        this.username = localStorage.getItem("username");        
        axios.post("http://localhost:9090/job_offer",{ 
            Publisher: this.username,
            Position: this.Position,
            JobDescription: this.JobDescription,
            Requirements: this.Requirements
            })
            .then (response => { 
                console.log("RESPONSE JE"+response.data)
                            new Swal({
                    title:"Uspesno",
                    type: "warning",
                    text:'Uspesno ste dodali ponudu za posao!',
                });     
                    this.$router.push({ name: "StartPageUser" });
            })       
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