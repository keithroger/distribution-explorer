export let distributions = [{
        name: "Normal",
        continous: true,
        params: ["Mu", "Sigma", "x"],
        args: ["0", "1", "0"],
        sliders: [
            {
                min: "-20",
                max: "20",
            }, {
                min: "1",
                max: "5",
            }, {
                min: "-20",
                max: "20",
            },
        ],
    }, {
        name: "Student's T",
        continous: true,
        modes: [
            {
                name: "Distribution",
                params: ["Sample Mean", "Sample Standard Deviation", "Degrees of Freedom"],
            }, {
                name: "PDF",
                params: ["Sample Mean", "Sample Standard Deviation", "Degrees of Freedom", "x"],
            }, {
                name: "CDF",
                params: ["Sample Mean", "Sample Standard Deviation", "Degrees of Freedom", "x"],
            }
        ],
    }, {
        name: "Chi-Squared",
        continous: true,
        modes: [
            {
                name: "Distribution",
                params: ["Degrees of Freedom"],
            }, {
                name: "PDF",
                params: ["Degrees of Freedom", "x"]
            }, {
                name: "CDF",
                params: ["Degrees of Freedom", "x"]
            }
        ],
    }, {
        name: "Gamma",
        continous: true,
        modes: [
            {
                name: "Distribution",
                params: ["Alpha", "Beta"],
            }, {
                name: "PDF",
                params: ["Alpha", "Beta", "x"]
            }, {
                name: "CDF",
                params: ["Alpha", "Beta", "x"]
            }
        ],
    }, {
        name: "Poisson",
        continous: false,
        modes: [
            {
                name: "Distribution",
                params: ["Lambda"],
            }, {
                name: "PDF",
                params: ["Lambda", "x"],
            }, {
                name: "CDF",
                params: ["Lambda", "x"],
            }
        ],
    }, {
        name: "Binomial",
        continous: false,
        modes: [
            {
                name: "Distribution",
                params: ["n", "p"],
            }, {
                name: "PDF",
                params: ["n", "p", "x"],
            }, {
                name: "CDF",
                params: ["n", "p", "x"],
            }
        ],
    }, {
        name: "Beta",
        continous: true,
        modes: [
            {
                name: "Distribution",
                params: ["Alpha", "Beta"],
            }, {
                name: "PDF",
                params: ["Alpha", "Beta", "x"],
            }, {
                name: "CDF",
                params: ["Alpha", "Beta", "x"],
            }
        ],
    }, {
        name: "Exponential",
        continous: true,
        modes: [
            {
                name: "Distribution",
                params: ["Lambda"],
            }, {
                name: "PDF",
                params: ["Lambda", "x"],
            }, {
                name: "CDF",
                params: ["Lambda", "x"],
            }
        ],
    }];