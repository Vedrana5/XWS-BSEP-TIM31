import { createRouter, createWebHistory } from "vue-router";

import StartPageView from "../views/startPage/StartPageView";
import LoginView from "../views/user/LoginView";
import RegistrationView from "../views/user/RegistrationView";
import FindPublicUser from "../views/user/FindPublicUser";
import UpdateInfoView from "../views/user/UpdateInfoView"
const routes = [
  {
    path: "/",
    name: "StartPageView",
    component: StartPageView,
  },
  {
  path: "/LoginView",
  name: "LoginView",
  component: LoginView,
  },
  {
    path: "/RegistrationView",
    name: "RegistrationView",
    component: RegistrationView,
    },
    {
      path: "/FindPublicUser",
      name: "FindPublicUser",
      component: FindPublicUser,
      },
      {
        path: "/UpdateInfoView",
        name: "UpdateInfoView",
        component: UpdateInfoView,
        },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
