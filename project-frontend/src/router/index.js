import { createRouter, createWebHistory } from "vue-router";

import StartPageView from "../views/startPage/StartPageView";
import LoginView from "../views/user/LoginView";
import RegistrationView from "../views/user/RegistrationView";
import FindPublicUser from "../views/user/FindPublicUser";
import UpdateInfoView from "../views/user/UpdateInfoView"
import StartPageUser from "../views/user/StartPageUser"
import ConfirmRegistration from "../views/user/ConfirmRegistration"
import ConfirmResetPassword from "../views/user/ConfirmResetPassword"
import ChangePassword from "../views/user/ChangePassword"
import ResetPassword from "../views/user/ResetPassword"
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
    path: "/ResetPassword",
    name: "ResetPassword",
    component: ResetPassword,
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
        {
        path: "/StartPageUser",
        name: "StartPageUser",
        component: StartPageUser,
      },
      {
        path: "/ChangePassword",
        name: "ChangePassword",
        component: ChangePassword,
      },
      {
        path: '/confirmRegistration/:confirmationToken/:userId',
        name: 'ConfirmRegistration',
        component: ConfirmRegistration
      },
      
      {
        path: '/confirmResetPassword/:id/:confirmationToken',
        name: 'ConfirmResetPassword',
        component: ConfirmResetPassword        
      }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
