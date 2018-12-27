package osu

/*
TODO: At some point, test using this data:

    pub fn get_test_data() -> Vec<(TimeLocation, i32)> {
        let test_data = vec![
            // uninherited timing points
            (TimeLocation::Relative(&TP, 0, Ratio::new(0, 1)), 12345), // no change from the measure at all
            (TimeLocation::Relative(&TP, 1, Ratio::new(0, 1)), 13545), // +1 measure (measure is 300ms, times 4 beats)
            (TimeLocation::Relative(&TP, 0, Ratio::new(1, 4)), 12645), // a single beat
            (TimeLocation::Relative(&TP, 0, Ratio::new(1, 2)), 12945), // half of a measure
            (TimeLocation::Relative(&TP, 0, Ratio::new(3, 4)), 13245), // 3 quarter notes
            // ok, on to inherited
            (TimeLocation::Relative(&ITP, 0, Ratio::new(0, 1)), 13545), // no change from the measure at all
            (TimeLocation::Relative(&ITP, 1, Ratio::new(0, 1)), 14745), // +1 measure, same as above
            (TimeLocation::Relative(&ITP, 0, Ratio::new(1, 4)), 13845), // a single beat
            (TimeLocation::Relative(&ITP, 0, Ratio::new(1, 2)), 14145), // half of a measure
            (TimeLocation::Relative(&ITP, 0, Ratio::new(3, 4)), 14445), // 3 quarter notes
        ];
        return test_data;
    }
*/
