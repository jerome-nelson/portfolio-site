module.exports = function (condition, message, args) {
    const { check, funcCheck } = args;
    let isConditionTrue = false;

    (() => {
      const checks = Array.isArray(condition) ? condition : [condition];
      checks.forEach(elem => {
        isConditionTrue = check
          ? elem === check
          : funcCheck ? funcCheck(elem) : false;
      });
    })();

    console.log(`${message}: ${isConditionTrue}`);
  };