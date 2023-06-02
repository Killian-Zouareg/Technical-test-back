package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Messages []Message

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	Messages = []Message{
		{701, "Inexcusable", "Meh"},
		{702, "Inexcusable", "Emacs"},
		{703, "Inexcusable", "Explosion"},
		{704, "Inexcusable", "Goto Fail"},
		{705, "Inexcusable", "I wrote the code and missed the necessary validation by an oversight (see 795)"},
		{706, "Inexcusable", "Delete Your Account"},
		{707, "Inexcusable", "Can't quit vi"},
		{710, "Novelty Implementations", "PHP"},
		{711, "Novelty Implementations", "Convenience Store"},
		{712, "Novelty Implementations", "NoSQL"},
		{718, "Novelty Implementations", "I am not a teapot"},
		{719, "Novelty Implementations", "Haskell"},
		{720, "Edge Cases", "Unpossible"},
		{721, "Edge Cases", "Known Unknowns"},
		{722, "Edge Cases", "Unknown Unknowns"},
		{723, "Edge Cases", "Tricky"},
		{724, "Edge Cases", "This line should be unreachable"},
		{725, "Edge Cases", "It works on my machine"},
		{726, "Edge Cases", "It's a feature, not a bug"},
		{727, "Edge Cases", "32 bits is plenty"},
		{728, "Edge Cases", "It works in my timezone"},
		{730, "Fucking", "Fucking npm"},
		{731, "Fucking", "Fucking Rubygems"},
		{732, "Fucking", "Fucking Unic💩de"},
		{733, "Fucking", "Fucking Deadlocks"},
		{734, "Fucking", "Fucking Deferreds"},
		{735, "Fucking", "Fucking IE"},
		{736, "Fucking", "Fucking Race Conditions"},
		{737, "Fucking", "FuckThreadsing"},
		{738, "Fucking", "Fucking Exactly-once Delivery"},
		{739, "Fucking", "Fucking Windows"},
		{738, "Fucking", "Fucking Exactly"},
		{739, "Fucking", "Fucking McAfee"},
		{750, "Syntax Errors", "Didn't bother to compile it"},
		{753, "Syntax Errors", "Syntax Error"},
		{754, "Syntax Errors", "Too many semi-colons"},
		{755, "Syntax Errors", "Not enough semi-colons"},
		{756, "Syntax Errors", "Insufficiently polite"},
		{757, "Syntax Errors", "Excessively polite"},
		{759, "Syntax Errors", "Unexpected T_PAAMAYIM_NEKUDOTAYIM"},
		{761, "Substance", "Hungover"},
		{762, "Substance", "Stoned"},
		{763, "Substance", "Under-Caffeinated"},
		{764, "Substance", "Over-Caffeinated"},
		{765, "Substance", "Railscamp"},
		{766, "Substance", "Sober"},
		{767, "Substance", "Drunk"},
		{768, "Substance", "Accidentally Took Sleeping Pills Instead Of Migraine Pills During Crunch Week"},
		{771, "Predictable Problems", "Cached for too long"},
		{772, "Predictable Problems", "Not cached"},
	}

	r.GET("/messages", getAllExcuses)
	r.POST("/AddExcuse", addExcuse)

	r.Run(":8080")
}
