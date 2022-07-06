<template>
    <div>
        <button @click="GoBack()">Go back</button>
        <label>Search by position</label>
        <input  placeholder="Type text" name="position" class="input-field" v-model="search"/>
        <button @click="Search(search)">Search</button>
        <button @click="SeeAllJobOffers()">See All</button>
        <div class="containerInfo" v-for="(jobOffer, index) in jobOffers" :key="index">
            <div class="tab-pane container active">
                <div class="row-boats" >
                    <div class="col-info">
                        <h4 style="width: 600px" class="text">
                            Publisher: {{jobOffer.Publisher}}
                        </h4>
                        <h4 style="width: 600px" class="text">Position: {{jobOffer.Position}}</h4>
                        <h4 style="width: 600px" class="text">
                            JobDescription: {{jobOffer.JobDescription}}
                        </h4>
                        <h4 style="width: 600px" class="text">
                            Requirement: 
                        </h4>                        
                        <div v-for="(Requirement, index) in Requirements " :key="index">
                            <label>{{Requirement}}</label>
                        </div>
                        <h4 style="width: 600px" class="text">
                        DatePosted: {{jobOffer.DatePosted}}
                        </h4>                        
                    </div>
                </div>
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
        Position:"",
        JobDescription:"",
        Requirements:[],
        Requirement:"",
        username:"",
        jobOffers:"",
        search:""
    };
  },


  methods: {
    async GoBack() {
        this.$router.push({ name: "StartPageUser" });
    },
    async Search(search) {
        this.search = search;
        console.log("Pretrazujes"+this.search);
        axios.get("http://localhost:9090/job_offer/"+this.search)
        .then (response => { 
        this.jobOffers = response.data.jobOffer;
        })        
    }, 
    async SeeAllJobOffers() {
        axios.get("http://localhost:9090/job_offer")
        .then (response => { 
        this.jobOffers = response.data.JobOffers;
        })
    },
  },
      async created() {
        await this.SeeAllJobOffers();
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