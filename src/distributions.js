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
        params: ["Sample Mean", "Sample Standard Deviation", "Degreees of Freedom", "x"],
        args: ["0", "1", "3", "0"],
        sliders: [
            {
                min: "-20",
                max: "20",
            }, {
                min: "1",
                max: "5",
            }, {
                min: "1",
                max: "50",
            }, {
                min: "-20",
                max: "20",
            },
        ],
    }, {
        name: "Chi-Squared",
        continous: true,
        params: ["Degrees of Freedom", "x"],
        args: ["4", "2"],
        sliders: [
            {
                min: "2",
                max: "20",
            }, {
                min: "0",
                max: "50",
            },
        ],
    }, {
        name: "Gamma",
        continous: true,
        params: ["Alpha", "Beta", "x"],
        args: ["5", "5", "5"],
        sliders: [
            {
                min: "1",
                max: "20",
            }, {
                min: "1",
                max: "20",
            }, {
                min: "0",
                max: "50",
            },
        ],
    }, {
        name: "Poisson",
        continous: false,
        params: ["Lambda", "x"],
        args: ["4", "2"],
        sliders: [
            {
                min: "1",
                max: "20",
            }, {
                min: "0",
                max: "50",
            },
        ],
    }, {
        name: "Binomial",
        continous: false,
        params: ["n", "p", "x"],
        args: ["10", "0.5", "6"],
        sliders: [
            {
                min: "1",
                max: "50",
            }, {
                min: "0.05",
                max: "0.95",
                step: "0.05",
            }, {
                min: "0",
                max: "20",
            },
        ],
    }, {
        name: "Beta",
        continous: true,
        params: ["Alpha", "Beta", "x"],
        args: ["1", "7", "0.5"],
        sliders: [
            {
                min: "1",
                max: "20",
            }, {
                min: "1",
                max: "20",
            }, {
                min: "0.05",
                max: "0.95",
                step: "0.05",
            },
        ],
    }, {
        name: "Exponential",
        continous: true,
        params: ["Lambda", "x"],
        args: ["1", "4"],
        sliders: [
            {
                min: "1",
                max: "20",
            }, {
                min: "0",
                max: "50",
            },
        ],
    }];