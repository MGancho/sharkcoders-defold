components {
  id: "doors"
  component: "/main/scripts/doors.script"
}
embedded_components {
  id: "right_door"
  type: "sprite"
  data: "default_animation: \"Right_door\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/atlas/doors.atlas\"\n"
  "}\n"
  ""
}
