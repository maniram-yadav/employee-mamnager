var EmployeeManager = artifacts.require("EmployeeManager");

module.exports = function(deployer) {
  // deployment steps
  deployer.deploy(EmployeeManager);
};