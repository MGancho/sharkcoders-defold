components {
  id: "doors"
  component: "/main/scripts/doors.script"
}
embedded_components {
  id: "left_door"
  type: "sprite"
  data: "default_animation: \"Left_door\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/atlas/doors.atlas\"\n"
  "}\n"
  ""
}
