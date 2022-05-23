let distributions = [
    {
        name: "Normal",
        formula: String.raw`f(x; \mu, \sigma) = \frac{1}{\sigma \sqrt{ 2 \pi }} \exp\left( - \frac{ (x - \mu)^2 } {2 \sigma^2}\right)`,
        continous: true,
        params: ["\\mu", "\\sigma", "x"],
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
        formula: String.raw`f(x; \hat{\mu}, \hat{\sigma}) = \frac{\Gamma((\nu+1)/2)}{\Gamma(\nu/2)\sqrt{\nu\pi\hat{\sigma}^2}} {\left(1+\dfrac{(x-\hat{\mu})^2}{\nu\hat{\sigma}^2}\right)}^{\!-(\nu+1)/2}`,
        continous: true,
        params: ["\\hat{\\mu}", "\\hat{\\sigma}", "\\nu", "x"],
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
        formula: String.raw`f(x; k) = \frac{e^{-x/2}x^{k/2-1}} {2^{k/2}\Gamma(k/2) }`,
        continous: true,
        params: ["k", "x"],
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
        formula: String.raw`f(x; \alpha, \beta) = \dfrac{x^{\alpha-1}e^{-x/\beta}}{\Gamma(\alpha)\beta^\alpha}`,
        continous: true,
        params: ["\\alpha", "\\beta", "x"],
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
        formula: String.raw`f(x; \lambda) = \dfrac{\lambda ^x}{x!}e^{-\lambda }`,
        continous: false,
        params: ["\\lambda", "x"],
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
        formula: String.raw`f(x;n,p)\binom{n}{x}p^x(1-p)^{n-x}`,
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
        formula: String.raw`f(x;\alpha,\beta) = \dfrac{\Gamma(\alpha + \beta)}{\Gamma(\alpha)\Gamma(\beta)}(1-x)^{\beta-1}x^{\alpha-1}`,
        continous: true,
        params: ["\\alpha", "\\beta", "x"],
        args: ["1", "7", "0.5"],
        sliders: [
            {
                min: "1.0",
                max: "20.0",
                step: "0.5",
            }, {
                min: "1.0",
                max: "20.0",
                step: "0.5",
            }, {
                min: "0.00",
                max: "1.00",
                step: "0.5",
            },
        ],
    }, {
        name: "Exponential",
        formula: String.raw`f(x;\lambda) = \lambda e^{-\lambda x}`,
        continous: true,
        params: ["\\lambda", "x"],
        args: ["1", "4"],
        sliders: [
            {
                min: "1",
                max: "10",
            }, {
                min: "0",
                max: "20",
            },
        ],
    }
];

export default distributions;