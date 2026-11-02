## staticconfig on zoltar

The zoltar design is in internal/zoltar/docs/design.md on the zoltar branch. The staticconfig port can start once its step 1 (the reshape) is done. Every staticconfig fixture must give the same results before and after the port.

### Example: staticconfig's route loop

```go
p := z.Predict(appConfigure, []zoltar.Value{
	zoltar.Unknown(appType, "the application"),
	configurer, // zoltar.Placeholder(configurerType)
})

routes := p.ValueOf(routesCall.Call.Args[0], routesCall)

for _, slice := range routes.Timelines() {
	elems, _ := slice.Elements() // ok: a single timeline
	speculative := !elems.IsComplete()

	for _, el := range elems.All() {
		_, indexKnown := el.Index()
		speculative := speculative || el.IsConditional() || !indexKnown

		for _, route := range el.Value().Timelines() {
			if route.HasBlindSpots() {
				// e.g. "ViaIntegration never returns" for ViaIntegration(nil)
				continue
			}
			o, _ := route.Origin()
			via := o.Prediction()               // the dogma.ViaX() prediction
			pos := via.CallSite().Pos()         // RoutedAt
			h := route.FieldByName("handler")   // already as of c.Routes
			hType, ok := h.DynamicType()        // e.g. *app.Handler
			...
		}
	}
}
```

### Example: should staticconfig follow this call?

```go
for _, arg := range target.Args {
	for x := range arg.Reachable() {
		if x.Equal(configurer) ||
			(x.HasBlindSpots() && typeMayHold(x.StaticType(), configurerType)) {
			// the configurer may be passed to this call: follow it
		}
	}
}
```

### Decided for staticconfig

- `dogma.ViaX(nil)`: zoltar sees that the call never returns. staticconfig
  reports "Configure() always panics" at the ViaX call, with a Related entry
  "nil handler passed to dogma.ViaX()". Recorded in the plan file.
- Recursion with differing arguments is cut off at the internal depth limit,
  not merged. Merging would change the memory model (one allocation standing
  for many variables) and is a separate research item.
