open Base

let read_file file = In_channel.with_open_text file In_channel.input_all
let current_step = ref 0
let basement_first_visited = ref 0

let rec count_floors input floor =
  match input with
  | [] -> floor
  | '(' :: rest -> count_floors rest (floor + 1)
  | ')' :: rest -> count_floors rest (floor - 1)
  | _ -> floor
;;

let rec find_first_basement_visit input floor =
  match input with
  | [] -> !basement_first_visited - 1
  | '(' :: rest ->
    Int.incr current_step;
    find_first_basement_visit rest (floor + 1)
  | ')' :: rest ->
    Int.incr current_step;
    if floor = -1 && !basement_first_visited = 0
    then basement_first_visited := !current_step;
    find_first_basement_visit rest (floor - 1)
  | _ -> !basement_first_visited - 1
;;

let part_one file =
  let floor = count_floors file 0 in
  Stdlib.print_endline ("Floor: " ^ Int.to_string floor)
;;

let part_two file =
  let floor = find_first_basement_visit file 0 in
  Stdlib.print_endline ("First basement visit: " ^ Int.to_string floor)
;;

let () =
  let file = read_file "bin/input.txt" |> String.to_list in
  part_one file;
  part_two file
;;
