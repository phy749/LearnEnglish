import AppHeader from "../../Components/Header";
import PageContent from "../../Components/PageContent";
import SideMenu from "../../Components/SideMenu";
// import "./managa.css";

function Admin() {
  return (
    <div className="App">
      <AppHeader />
      <div className="SideMenuAndPageContent">
        <SideMenu/>
        <PageContent/>
      </div>
    </div>
  );
}
export default Admin;

