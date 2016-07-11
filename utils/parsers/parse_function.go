package parsers

var (
	equality = "="
)

// Function will take a string argument and return a function, else error
// Example format for Function Parser
// "f(x,y)=(Sqrt(Sin(x)^2+Cos(y)^2)+6*(1-5i))*`[1 2 3]"
