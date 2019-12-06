module.exports = {
    main: async function (event, context) {
        console.log("inside lambda");
        return "Hello World";
    }
}